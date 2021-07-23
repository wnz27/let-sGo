/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/7/24 01:31 7月
 **/
package main

import (
	"context"
	pb "fzkprac/prac_grpc/learn_20210724/grpc_demo/route"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"log"
	"math"
	"net"
)

type routeGuideServer struct {
	features []*pb.Feature  // 相当于假db
	pb.UnimplementedRouteGuidServer
}

// check if a point is inside a rectangle
func inRange(point *pb.Point, rect *pb.Rectangle) bool {
	left := math.Min(float64(rect.Lo.Longitude), float64(rect.Li.Longitude))
	right := math.Max(float64(rect.Lo.Longitude), float64(rect.Li.Longitude))
	top := math.Max(float64(rect.Lo.Latitude), float64(rect.Li.Latitude))
	bottom := math.Min(float64(rect.Lo.Latitude), float64(rect.Li.Latitude))

	if float64(point.Longitude) >= left &&
		float64(point.Longitude) <= right &&
		float64(point.Latitude) >= bottom &&
		float64(point.Latitude) <= top {
		return true
	}
	return false
}

func (r *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, f := range r.features {
		if proto.Equal(f.Location, point) {
			return f, nil
		}
	}
	return nil, nil
}

func (r *routeGuideServer) ListFeatures(rectangle *pb.Rectangle, server pb.RouteGuid_ListFeaturesServer) error {
	for _, f := range r.features {
		if inRange(f.Location, rectangle) {
			if err := server.Send(f); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *routeGuideServer) RecordRoute(server pb.RouteGuid_RecordRouteServer) error {
	return nil
}

func (r *routeGuideServer) Recommend(server pb.RouteGuid_RecommendServer) error {
	return nil
}

//func (r routeGuideServer) mustEmbedUnimplementedRouteGuidServer() {
//	print()
//}

func newServer() *routeGuideServer {
	return &routeGuideServer{
		features: []*pb.Feature{
			{Name: "上海交通大学闵行校区 上海市闵行区东川路800号", Location: &pb.Point{
				Latitude:  310235000,
				Longitude: 121437403,
			}},
			{Name: "复旦大学 上海市杨浦区五角场邯郸路220号", Location: &pb.Point{
				Latitude:  312978870,
				Longitude: 121503457,
			}},
			{Name: "华东理工大学 上海市徐汇区梅陇路130号", Location: &pb.Point{
				Latitude:  311416130,
				Longitude: 121424904,
			}},
		},
	}
}

func main() {
	con, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	// 把服务注册给 grpc server
	pb.RegisterRouteGuidServer(grpcServer, newServer())
	log.Fatal(
		grpcServer.Serve(con))
}

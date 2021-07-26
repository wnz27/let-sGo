/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/7/24 01:31 7月
 **/
package main

import (
	"context"
	"fmt"
	pb "fzkprac/prac_grpc/learn_20210724/grpc_demo/route"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"math"
	"net"
	"time"
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

func toRadians(num float64) float64 {
	return num * math.Pi / float64(180)
}
// 给地球 俩点计算距离 结果单位 米 m
// calcDistance calculates the distance between two points using the "haversine" formula.
// The formula is based on http://mathforum.org/library/drmath/view/51879.html.
func calcDistance(p1 *pb.Point, p2 *pb.Point) int32 {
	const CordFactor float64 = 1e7
	const R = float64(6371000) // earth radius in metres
	lat1 := toRadians(float64(p1.Latitude) / CordFactor)
	lat2 := toRadians(float64(p2.Latitude) / CordFactor)
	lng1 := toRadians(float64(p1.Longitude) / CordFactor)
	lng2 := toRadians(float64(p2.Longitude) / CordFactor)
	dlat := lat2 - lat1
	dlng := lng2 - lng1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	return int32(distance)
}

func (r *routeGuideServer) RecordRoute(server pb.RouteGuid_RecordRouteServer) error {
	startTime := time.Now()
	var pointCount, distance int32
	var prevPoint *pb.Point
	for {
		point, err := server.Recv()
		if err == io.EOF {
			// conclude a route summary
			endTime := time.Now()
			return server.SendAndClose(&pb.RouteSummary{
				PointCount: pointCount,
				Distance: distance,
				ElapsedTime: int32(endTime.Sub(startTime).Seconds()),
			})
		}
		if err != nil {
			return err
		}
		pointCount ++
		if prevPoint != nil {
			distance += calcDistance(prevPoint, point)
		}
		prevPoint = point
	}
}

func (r *routeGuideServer) recommendOnce(req *pb.RecommendationRequest) (*pb.Feature, error) {
	var nearest, farthest *pb.Feature
	var nearestDistance, farthestDistance int32

	for _, feature := range r.features {
		distance := calcDistance(feature.Location, req.Point)

		if nearest == nil || distance < nearestDistance {
			nearestDistance = distance
			nearest = feature
		}

		if farthest == nil || distance > farthestDistance {
			farthestDistance = distance
			farthest = feature
		}
	}

	if req.Mode == pb.RecommendationMode_GetFarthest {
		return farthest, nil
	} else {
		return nearest, nil
	}
}

func (r *routeGuideServer) Recommend(server pb.RouteGuid_RecommendServer) error {
	for {
		req, err := server.Recv()
		fmt.Println("0000000000000000------>1", err)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		recommendedFeature, err := r.recommendOnce(req)
		if err != nil {
			return err
		}
		return server.Send(recommendedFeature)
	}
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

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
	"log"
	"net"
)

type routeGuideServer struct {
	pb.UnimplementedRouteGuidServer
}

func (r routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	return nil, nil
}

func (r routeGuideServer) ListFeatures(rectangle *pb.Rectangle, server pb.RouteGuid_ListFeaturesServer) error {
	return nil
}

func (r routeGuideServer) RecordRoute(server pb.RouteGuid_RecordRouteServer) error {
	return nil
}

func (r routeGuideServer) Recommend(server pb.RouteGuid_RecommendServer) error {
	return nil
}

//func (r routeGuideServer) mustEmbedUnimplementedRouteGuidServer() {
//	print()
//}

func newServer() *routeGuideServer {
	return &routeGuideServer{}
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

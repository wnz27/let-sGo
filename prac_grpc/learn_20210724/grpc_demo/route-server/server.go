/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/7/24 01:31 7月
 **/
package main

import (
	"google.golang.org/grpc"
	"net"
	pb "fzkprac/prac_grpc/learn_20210724/grpc_demo/route"
)

type RouteGuideServer {

}


func newServer() *pb.RouteGuidServer {

}

func main() {
	con, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	// 把服务注册给 grpc server
	pb.RegisterRouteGuidServer(grpcServer, )
}

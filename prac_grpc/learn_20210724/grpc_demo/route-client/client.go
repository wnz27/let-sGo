/**
 * @project let-sGo
 * @Author 27
 * @Description 
 * @Date 2021/7/24 01:31 7月
 **/
package main

import (
	"context"
	"fmt"
	pb "fzkprac/prac_grpc/learn_20210724/grpc_demo/route"
	"google.golang.org/grpc"
	"io"
)

func runFirst(c pb.RouteGuidClient) {
	feature, err := c.GetFeature(context.Background(), &pb.Point{
		Latitude:  310235000,
		Longitude: 121437403,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(feature)
}

func runSecond(c pb.RouteGuidClient) {
	serverStream, err := c.ListFeatures(context.Background(), &pb.Rectangle{
		Lo: &pb.Point{Latitude: 313374060, Longitude: 121358540},
		Li: &pb.Point{Latitude: 311034130, Longitude: 121598790},
	})
	if err != nil {
		panic(err)
	}
	count := 1
	for {
		feature, err := serverStream.Recv()
		if err == io.EOF {   // 流结束也会有一个err, io.EOF 是 正常的表现，所以只是作为结束的标识
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("count: %d, %v \n", count, feature)
		count ++
	}
}

func main() {
	con, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		panic(err)
	}
	defer con.Close()

	client := pb.NewRouteGuidClient(con)
	//runFirst(client)
	runSecond(client)
}


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

func main() {
	con, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		panic(err)
	}
	defer con.Close()

	client := pb.NewRouteGuidClient(con)
	runFirst(client)

}


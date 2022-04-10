/**
 * @project let-sGo
 * @Author 27
 * @Description 
 * @Date 2021/7/24 01:31 7月
 **/
package main

import (
	"bufio"
	"context"
	"fmt"
	pb "fzkprac/prac_grpc/learn_20210724/grpc_demo/route"
	"google.golang.org/grpc"
	"io"
	"os"
	"log"
	"time"
)

// unary
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

// server side stream
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

// client side stream
func runThird(c pb.RouteGuidClient) {
	// dummy data 假设客户上传这些数据
	points := []*pb.Point{
		{Latitude: 313374060, Longitude: 121358540},
		{Latitude: 311034130, Longitude: 121598790},
		{Latitude: 310235000, Longitude: 121437403},
	}

	clientStream, err := c.RecordRoute(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, point := range points {
		if err := clientStream.Send(point); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
	summary, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(summary)
}

// todo 从标准流里读一个int 存在target里面
func readIntFromCommendLine(reader *bufio.Reader, target *int32) {
	_, err := fmt.Fscanf(reader, "%d\n", target)
	if err != nil {
		log.Fatalln("Cannot scan", err)
	}
}


// bidirectional streaming
func runForth(c pb.RouteGuidClient) {
	serverStream, err := c.Recommend(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	// this goroutine listen to the server stream
	go func() {
		feature, err2 := serverStream.Recv()  // 从server接收值
		if err2 != nil {
			log.Fatalln(err2)
		}
		fmt.Println("Recommended: ", feature)
	}()

	reader := bufio.NewReader(os.Stdin)  // 做一个从标准输入读值的reader

	// 做一个和命令行交互的东西发送请求
	// 输入经纬度向服务器询问最近或者最远的Feature, 利用io reader
	for {
		request := pb.RecommendationRequest{
			Point: new(pb.Point),
		}
		var mode int32
		fmt.Print("Enter Recommendation Mode (0 for farthest, 1 for the nearest)")
		readIntFromCommendLine(reader, &mode)
		fmt.Print("Enter Latitude:")
		readIntFromCommendLine(reader, &request.Point.Latitude)
		fmt.Print("Enter Longitude:")
		readIntFromCommendLine(reader, &request.Point.Longitude)
		request.Mode = pb.RecommendationMode(mode)
		// 再发回给server reques1t
		if err2 := serverStream.Send(&request); err2 != nil {
			log.Fatalln(err2)
		}
		time.Sleep(100 * time.Millisecond)
	}

}

func main() {
	con, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("client can not dail server")
	}
	defer con.Close()

	client := pb.NewRouteGuidClient(con)
	//runFirst(client)
	//runSecond(client)
	//runThird(client)
	runForth(client)
}


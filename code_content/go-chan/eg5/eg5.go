/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/5/1 17:03 5月
 **/
package main

import (
	"errors"
	"fmt"
	"time"
)

/*
模拟远程过程调用（RPC）
服务器开发中会使用RPC（Remote Procedure Call，远程过程调用）简化进程间通信的过程。
RPC能有效地封装通信过程，让远程的数据收发通信过程看起来就像本地的函数调用一样。

本例中，使用通道代替Socket实现RPC的过程。客户端与服务器运行在同一个进程，服务器和客户端在两个goroutine中运行。
 */

/*
1．客户端请求和接收封装
下面的代码封装了向服务器请求数据，等待服务器返回数据，如果请求方超时，该函数还会处理超时逻辑
 */

func RPCClient(ch chan string, req string) (string, error) {
	// 向服务器发送请求
	ch <- req

	// 等待服务器返回
	select {
	case ack := <-ch:  // 接收到服务器返回数据
		return ack, nil
	case <- time.After(time.Second):  //  超时
		return "time problem", errors.New("Time out")
	}

}
// 2．服务器接收和反馈数据
//服务器接收到客户端的任意数据后，先打印再通过通道返回给客户端一个固定字符串，表示服务器已经收到请求。
func RPCServer (ch chan string) {
	for {
		// 接收客户端请求
		data := <-ch

		// 打印接收到的数据
		fmt.Println("server received:", data)

		// 客户端反馈已收到
		ch <- "roger"
	}
}

//3．模拟超时
func TimeoutRPCServer (ch chan string) {
	for {
		// 接收客户端请求
		data := <-ch

		// 打印接收到的数据
		fmt.Println("server received:", data)

		// 手动超时2秒
		time.Sleep(time.Second * 2)

		// 客户端反馈已收到
		ch <- "roger"
	}
}


func main(){

	ch := make(chan string)

	go RPCServer(ch)

	// 非超时场景
	resp, err := RPCClient(ch, "hi") ; if err != nil {
		panic("123")
	}

	fmt.Println("client received:", resp)


	ch2 := make(chan string)
	// 超时场景
	go TimeoutRPCServer(ch2)

	resp2, err2 := RPCClient(ch2, "hi") ; if err2 != nil {
		fmt.Println("res:", resp2, "err:", err2)
	}
	fmt.Println("res:", resp2, "err:", err2)


}

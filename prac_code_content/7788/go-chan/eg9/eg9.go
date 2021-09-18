/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/1 8:19 下午
* Description:
 */
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Telnet回音服务器——TCP服务器的基本结构

/*
Go语言中可以根据实际会话数量创建多个goroutine，并自动的调度它们的处理。

telnet 服务器处理
 */
// 1．接受连接
func server(address string, exitChan chan int){
	// 根据给定地址进行侦听
	l, err := net.Listen("tcp", address)

	// 如果侦听发生错误，打印错误并退出
	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}

	// 打印侦听地址，表示侦听成功
	fmt.Println("listen: ", address)

	// 延迟关闭侦听器
	defer l.Close()

	// 侦听循环
	for {
		// 新链接没有来的时候 Accept是阻塞的
		conn, err := l.Accept()
		// 发生任何的侦听错误，打印错误并退出服务器
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		// 根据链接开启会话，这个过程需要并行执行
		go handleSession(conn, exitChan)
	}
	fmt.Println("asdfasadf")
}

// 2．会话处理
// 回音服务器需要将收到的有效数据通过Socket发送回去。
func handleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("session start")

	// 创建一个网络连接数据的读取器
	reader := bufio.NewReader(conn)

	// 接收数据的循环
	for {
		// 读取字符串，碰到回车返回
		str, err := reader.ReadString('\n')

		// 数据读取正确
		if err == nil {
			// 去掉字符串尾部的回车
			str = strings.TrimSpace(str)

			// 处理Telnet 指令
			if !processTelnetCommend(str, exitChan){
				conn.Close()
				break
			}

			// Echo 逻辑 发送什么数据，原样返回
			conn.Write([]byte(str + "\r\n"))
		} else {
			// 发生错误
			fmt.Println("session closed")
			conn.Close()
			break
		}
	}
}

// 3．Telnet命令处理
// 在下例中，定义了以下两个特殊控制指令，用以实现一些功能：
// ● 输入“@close”退出当前连接会话。
// ● 输入“@shutdown”终止服务器运行。
func processTelnetCommend(str string, exitChan chan int) bool {
	// @close 指令表示终止本次会话
	if strings.HasPrefix(str, "@close") {
		fmt.Println("session closed")
		// 告诉外部需要断开连接
		return false

		// @shutdown指令表示终止服务进程
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("server shutdown")

		// 往通道中写入0， 阻塞等待接收方处理
		exitChan <- 0

		// 告诉外部需要断开连接
		return false
	}

	// 打印输入的字符串
	fmt.Println("inner handle", str)
	return true
}


// 暂时跑不起来~~~ 输入telnet 会卡住
func main(){
	// 4．程序入口代码      Telnet 回音处理主流程
	// 创建一个程序结束码的通道
	exitChan := make(chan int)

	// 将服务器并发运行
	go server("127.0.0.1:7001", exitChan)

	// 通道阻塞，等待接收返回值
	code := <- exitChan

	// 标记程序返回值并退出
	os.Exit(code)

}

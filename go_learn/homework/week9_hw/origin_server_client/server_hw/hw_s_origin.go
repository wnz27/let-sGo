/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/6/27 15:07 6月
 **/
package main

import (
	"bufio"
	"io"
	"net"
	"fmt"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client发来的数据：", recvStr)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:40000")
	if err != nil {
		fmt.Println("listen tcp failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}

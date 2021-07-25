/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/7/25 15:57 7月
 **/
package main

import (
	"log"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 4096)
	conn.Read(buf)
	conn.Write([]byte("pong: "))
	conn.Write(buf)
}

func main() {
	server, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("err:%v", err)
	}
	for {
		c, err := server.Accept()
		if err != nil {
			log.Fatalf("err:%v", err)
		}
		go handleConn(c)
	}
}

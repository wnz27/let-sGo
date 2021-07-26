/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/7/25 15:57 7月
 **/
package main

import (
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("hello world\n"))
	io.Copy(os.Stdout, conn)
}

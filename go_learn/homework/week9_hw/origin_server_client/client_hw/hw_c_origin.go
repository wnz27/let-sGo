/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/6/27 15:07 6月
 **/
package main

import (
	"net"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:40000")
	if err != nil {
		fmt.Println("Connect tcp failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hi. How are you?`
		conn.Write([]byte(msg))
	}
}

/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/12 02:36 8月
 **/
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var stdoutBuff bytes.Buffer  // 创建一个内存缓冲区，以帮助减少输出的不确定性。
	// 它没有给我们任何保证，但它比直接写stdout 要快一些

	defer stdoutBuff.WriteTo(os.Stdout)  // 确保程序退出之前缓冲区内容需要被写入到stdout

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i:=0; i < 5; i ++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v. \n", integer)
	}
}

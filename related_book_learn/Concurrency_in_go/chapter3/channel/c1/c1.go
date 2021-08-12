/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/11 03:21 8月
 **/
package main

import "fmt"

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "hello"
	}()
	salutation, ok := <- stringStream
	fmt.Printf("(ok: %v): %v\n", ok, salutation)
	// 输出如下：
	// (ok: true): hello

	intStream := make(chan int)
	close(intStream)
	integer, ok := <- intStream
	fmt.Printf("(ok? %v): %v", ok, integer)

}

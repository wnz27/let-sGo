/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/11 01:45 8月
 **/
package main

import "fmt"

func main() {
	//var dChan chan<- interface{}
	//dChan = make(chan<- interface{})
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"  // 我们将字符串文本传递到 stringStream channel
	}()
	fmt.Println(<-stringStream)  // 我们读取channel的字符串字面量并将其打印到 stdout

	// 往只读写入，从只写读出，编译会报错
	//writeStream := make(chan<- interface{})
	//readStream := make(<-chan interface{})

	//<-writeStream
	//readStream <- struct{}{}
	/*
	invalid operation: <-writeStream(receive from send-only type chan<- interface{})
	invalid operation: readStream <- struct {} literal (send to receive-only type <-chan interface{})
	 */
}

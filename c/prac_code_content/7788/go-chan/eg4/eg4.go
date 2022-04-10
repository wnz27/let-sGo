/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/5/1 16:53 5月
 **/
package main

import "fmt"

func singleOrient() {
	// 只能 （向ch） 发送值 的 通道
	chSendOnly := make(chan<- int)
	chSendOnly <- 9

	// 只能 （从ch） 接收值 的 通道
	chRecvOnly := make(<-chan int)
	<- chRecvOnly
}

/*
无缓冲通道保证收发过程同步，无缓冲收发过程类似于快递员给你电话让你下楼取快递，整个收发过程同步，你和快递员不见不散。
但是快递员把快递放入快递柜中，并通知用户来取，快递员和用户就成了异步收发过程，效率可以有明显的提升，带缓冲的通道就是这样的一个快递柜。
*/

func main() {
	ch := make(chan int, 3)

	fmt.Println(len(ch))

	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println(len(ch))
}

/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/5/1 10:44 5月
 **/
package main

import "fmt"


// <-ch  执行这个语句会发送阻塞，直到接收数据但是接收到的数据会被忽略。
//这个方式实际上只是通过通道在goroutine间阻塞收发实现并发同步。

func main() {
	// 构建一个通道
	ch := make(chan int)

	// 开启一个并发匿名函数
	go func() {
		fmt.Println("start goroutine")
		// 通过通道通知main 的 goroutine
		ch <- 0
		fmt.Println("exit goroutine")
	}()

	fmt.Println("wait goroutine")

	// 等待匿名goroutine
	<-ch

	fmt.Println("all done")
}

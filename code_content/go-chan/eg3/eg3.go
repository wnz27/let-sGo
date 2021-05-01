/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/5/1 11:06 5月
 **/
package main

import "fmt"

// 并发打印

func printer(c chan int) {
	// 开始无限循环等待数据
	for {
		// 从channel 中获取一个数据
		data := <-c

		// 将0视为数据结束
		if data == 0 {
			break
		}

		// 打印数据
		fmt.Println(data)
	}

	// 通知main 已经结束循环(i got it)
	c <-0
}

func main(){
	ch := make(chan int)

	// 并发执行printer 传入channel， 消费者
	go printer(ch)

	// 生产者
	for i := 1; i <= 10; i ++ {
		// 将数据通过channel传给printer
		ch <- i
	}

	// 通知并发的printer结束循环(没数据了)
	ch <- 0

	// 等待printer 结束
	<-ch

}

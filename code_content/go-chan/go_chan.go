/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/4/29 1:40 下午
* Description:
 */
package main

import (
	"fmt"
	"time"
)

func slowFunc(c chan<- string){
	time.Sleep(time.Second * 2)
	c <- "slowFunc() finished!"
}

func receivers(c <-chan string){
	for msg := range c {
		fmt.Println(msg)
	}
}


func pinger(c chan<- string) time.Time{
	t := time.NewTicker(300 * time.Millisecond)
	for {
		c <- "Ping"
		<- t.C
	}
}

func sender(c chan<- string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "I'm sending a message"
		<- t.C
	}
}


func main() {
	//funcs.PipelineWork()
	fmt.Println(" ---------------------- ----  ---------------------- ")
	//funcs.Pipeline2()
	fmt.Println(" ---------------------- ----  ---------------------- ")
	//funcs.PipelineWork2()
	fmt.Println(" ---------------------- base prac  ---------------------- ")
	//go funcs.Tg()  // 不显示为啥？
	//time.Sleep(time.Second * 1)

	// 1、
	//a := funcs.OutFunc()
	//print(a)
	// 2、
	//c := make(chan string)
	//go slowFunc(c)
	//
	//msg := <- c
	//fmt.Println(msg)

	// 有缓冲
	//messages := make(chan string, 2)
	//messages <- "hello"
	//messages <- "world"
	//close(messages)
	//fmt.Println("Pushed two messages onto Channel with no receivers")
	//time.Sleep(time.Second * 1)
	//go receivers(messages)
	//fmt.Println("lslslsls")
	// 输出顺序不一定的


	// 阻塞和流程控制, TODO 疑问 这里为什么只打印除了Ping, 解释：因为msg只会接收一个，然后打印然后退出
	// TODO 因为无缓冲通道满的话就是阻塞的，pinger所谓的死循环不存在
	//pC := make(chan string)
	//go pinger(pC)
	//msg := <- pC
	//fmt.Println(msg)

	// 这里只会有3次打印，如果这里也写死循环那么会无限输出
	//pC := make(chan string)
	//go pinger(pC)
	//for i := 0; i< 3 ; i++ {
	//	msg := <-pC
	//	fmt.Println(msg)
	//}

	/*
	<-位于关键字左边时，表示通道在函数内是只读的。位于右边表示通道在函数内是只写的（我理解是只能往里写。）
	没有指定<- 时， 通道可读可写
	可以隐式转换，可读可写可以隐式转换为另外两个，但另外两个不能转换为单个功能的。
	 */

	// 接收停止信号
	//messages1 := make(chan string)
	//stop := make(chan bool)
	//go sender(messages1)
	//go func() {
	//	time.Sleep(time.Second * 3)
	//	fmt.Println("Time is up!")
	//	stop <- true
	//}()
	//
	//for {
	//	select {
	//	case <- stop:
	//		return
	//	case msg := <-messages1:
	//		fmt.Println(msg)
	//	}
	//}

	// fatal error: all goroutines are asleep - deadlock!
	//ch1 := make(chan int)
	//ch1 <- 0

}


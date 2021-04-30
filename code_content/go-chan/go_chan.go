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
	pC := make(chan string)
	go pinger(pC)
	for i := 0; i< 3 ; i++ {
		msg := <-pC
		fmt.Println(msg)
	}


}


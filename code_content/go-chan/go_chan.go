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



}


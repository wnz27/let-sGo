/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/1 7:20 下午
* Description:
 */
package main

import (
	"fmt"
	"time"
)

// 示例：使用通道响应计时器的事件

func main() {
	// 1．一段时间之后（time.After）
	// 声明一个退出用的通道
	exit := make(chan int)
	//打印开始
	fmt.Println("start")
	// 过一秒之后 调用匿名函数
	time.AfterFunc(time.Second, func() {
		// 1秒后打印
		fmt.Println("one second after")

		// 通知main() goroutine已经结束
		exit <- 0
	})
	// 等待结束
	<-exit
}

/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/1 7:46 下午
* Description:
 */
package main

import (
	"fmt"
	"time"
)

/*
2．定点计时

计时器（Timer）的原理和倒计时闹钟类似，都是给定多少时间后触发。

打点器（Ticker）的原理和钟表类似，钟表每到整点就会触发。

这两种方法创建后会返回time.Ticker对象和time.Timer对象，里面通过一个C成员，类型是只能接收的时间通道（<-chan Time），
使用这个通道就可以获得时间触发的通知。

下面代码创建一个打点器，每500毫秒触发一次；创建一个计时器，2秒后触发，只触发一次。
 */
func main() {
	// 创建一个打点器，每500 毫秒触发一次
	ticker := time.NewTicker(time.Millisecond * 500)

	// 创建一个定时器，2秒后触发
	stopper := time.NewTimer(time.Second * 2)

	//声明计数变量
	var i int

	// 不断检查通道情况
	for {
		// 多路复用通道
		select {
		case <- stopper.C:    // 计时器到了
			fmt.Println("stop")
			goto StopHere
		case <- ticker.C:  // 打点器触发
			// 记录触发了多少次
			i ++
			fmt.Println("tick", i)
		}
	}

	// 退出的标签 使用goto跳转
	StopHere:
		fmt.Println("done")
}

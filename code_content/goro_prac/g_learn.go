/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/6/23 11:37 下午
* Description:
 */

/*
实现控制并发的方式，大致可分成以下三类：

全局共享变量

channel通信

Context包
 */

/*
todo 1、这是最简单的实现控制并发的方式，实现步骤是：
	- 声明一个全局变量；
	- 所有子goroutine共享这个变量，并不断轮询这个变量检查是否有更新；
	- 在主进程中变更该全局变量；
	- 子goroutine检测到全局变量更新，执行相应的逻辑。
*/
package main

import (
	"fmt"
	"time"
)

func type1() {
	running := true

	f := func() {
		for running {
			fmt.Println("sub proc running...")
			time.Sleep(1 * time.Second)
		}
		fmt.Println("sub proc exit")
	}

	go f()
	go f()
	go f()

	time.Sleep(2 * time.Second)

	running = false

	time.Sleep(3 * time.Second)

	fmt.Println("main proc exit")
}
/*
todo 全局变量的优势是简单方便，不需要过多繁杂的操作，通过一个变量就可以控制所有子goroutine的开始和结束；
 缺点是功能有限，由于架构所致，该全局变量只能是多读一写，否则会出现数据同步问题，当然也可以通过给全局变量加锁来解决这个问题，
 但那就增加了复杂度，另外这种方式不适合用于子goroutine间的通信，因为全局变量可以传递的信息很小；还有就是主进程无法等待所有子goroutine退出，
 因为这种方式只能是单向通知，所以这种方法只适用于非常简单的逻辑且并发量不太大的场景，一旦逻辑稍微复杂一点，这种方法就有点捉襟见肘。

*/


/*
TODO channel通信
 Channel是Go中的一个核心类型，你可以把它看成一个管道，通过它并发核心单元就可以发送或者接收数据进行通讯(communication)。
	要想理解 channel 要先知道 CSP 模型：
	CSP 是 Communicating Sequential Process 的简称，中文可以叫做通信顺序进程，是一种并发编程模型，由 Tony Hoare 于 1977 年提出。
	简单来说，CSP 模型由并发执行的实体（线程或者进程）所组成，实体之间通过发送消息进行通信，这里发送消息时使用的就是通道，或者叫 channel。
	CSP 模型的关键是关注 channel，而不关注发送消息的实体。Go 语言实现了 CSP 部分理论，goroutine 对应 CSP 中并发执行的实体，channel 也就对应着 CSP 中的 channel。
	也就是说，CSP 描述这样一种并发模型：多个Process 使用一个 Channel 进行通信, 这个 Channel 连结的 Process 通常是匿名的，消息传递通常是同步的（有别于 Actor Model）。

 */
func type2() {
	fmt.Println(1)
}

func main() {



}



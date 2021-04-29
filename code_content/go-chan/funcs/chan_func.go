/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/4/29 1:40 下午
* Description:
 */

package funcs

import "fmt"

func TChan() {
	ch := make(chan int)
	//send
	ch <- 1
	// 赋值接收表达式
	x := <-ch
	// 接收语句， 丢弃结果
	<-ch
	fmt.Println("TChan", x)
	close(ch)
}

/*
关闭后的发送操作将崩溃宕机。在一个已经关闭的通道上进行接收操作将获取所有已经发送的值，直到通道为空。（意味着生产中这样会导致重复消费？）
这时任何接收操作会立即完成，同时获取到一个通道元素类型对应的零值。

 */

func PipelineWork() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; x<100 ; x ++ {
		//for x := 0; ; x ++ {
			naturals <- x
		}
		close(naturals)
	}()

	// squarer
	go func() {
		for {
			x, ok := <- naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	//for x := 0; x < 10 ; x ++{
	for x := 0; ; x ++{
		fmt.Println(<-squares)
	}
}
	


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

func counter(out chan<- int){
	for x := 0; x < 100; x ++ {
		out <- x
	}
	close(out)
}

func square(out chan<- int, in <-chan int) {
	for v := range in{
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}

}

func Pipeline2(){
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go square(squares, naturals)
	printer(squares)
}

func PipelineWork2() {
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
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// printer
	for x := range squares{
		fmt.Println(x)
	}
}

// 该方法有误
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
	// printer
	//for x := 0; x < 10 ; x ++{
	for x := 0; ; x ++{
		fmt.Println(<-squares)
	}
}
	


/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/12 00:59 8月
 **/
package main

import "fmt"

func main()  {
	intStream := make(chan int)
	go func() {
		defer close(intStream)  // 确保goroutine 退出之前channel是关闭的。这是一个常见的模式。
		for i := 1; i <= 5; i ++ {
			intStream <- i
		}
	}()
	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}
}




/**
 * @project let-sGo
 * @Author 27
 * @Description
 * @Date 2021/8/8 11:41 8月
 **/
package main

import (
	"fmt"
)

func main() {
	chanOwner := func() <- chan int{
		/*
		函数范围内实例化channel 这将结果写入channel的处理范围约束在它西面定义的闭包中
		换句话说，它包含了这个channel 的写入处理，以防止其他goroutine写入他
		 */
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 6; i++{
				results <- i
				//time.Sleep(time.Second)
			}
		}()
		return results
	}
	consumer := func(results <-chan int) {
		/*
		收到一个int channel 的只读副本。通过声明我们要求的唯一用法是读取访问，
		我们将channel 内的使用约束为只读
		 */
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	/*
	这里收到的是channel 的读处理，能够将他传递给消费者，消费者只能从中读取信息。
	这一次将main goroutine 约束在channel 的只读视图中。
	 */
	results := chanOwner()
	consumer(results)
}

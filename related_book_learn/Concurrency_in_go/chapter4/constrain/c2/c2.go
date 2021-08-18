/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/14 01:44 8月
 **/
package main

import "fmt"

func main() {
	// 在chanOwner函数的词法范围内实例化channel。这将结果写入channel的处理的范围约束在它的下面定义的闭包中。
	// 换句话说，它包含了这个chanel的写入处理，以防其他goroutine写入它。
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i ++ {
				results <- i
			}
		}()
		return results
	}

	// 收到一个int channel的只读副本。通过声明我们要求的唯一用法是读取访问，我们将channel内的使用约束为只读。
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
	}
	// 收到了channel的读处理，能够将它传递给消费者，消费者只能从中读取信息。
	// 这又一次将main goroutine 约束在channel的只读视图中。
	results := chanOwner()
	consumer(results)
}

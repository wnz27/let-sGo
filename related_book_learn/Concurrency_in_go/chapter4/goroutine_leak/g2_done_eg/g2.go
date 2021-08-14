/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/15 02:29 8月
 **/
package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface {} {  // 我们将完成的channel 传递给doWork函数。作为惯例，这个channel是第一个参数。
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					// 做一些有意思的事情
					fmt.Println(s)
				// 在下面这一行上，我们看到了在实际编程中无处不在的select模式。
				// 我们的一个案例陈述是检查我们的done channel是否已经发出信号。
				// 如果有的话，我们从goroutine返回
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {  // 在这里我们创建另一个 goroutine，如果超过 1s 就会取消doWork 中产生的 goroutine。
		// 在 1s 之后取消本操作
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-terminated  // 这就是我们加入从 main goroutine 的 doWork 中产生的 goroutine 的地方。 （这里也会阻塞等待close）
	fmt.Println("Done")
}

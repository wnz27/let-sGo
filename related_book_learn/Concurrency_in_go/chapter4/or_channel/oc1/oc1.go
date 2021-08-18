/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/15 04:07 8月
 **/
package main

import (
	"fmt"
	"time"
)

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	// 这就是我们的函数，它将一堆包含可变切片的channel 整合成一个单一的channel
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		// 因为这是一个循环函数，我们必须设置结束条件。首先，如果可变切片是空的，我们只返回一个空 channel。
		// 这里与我们传递进一个空channel 是相通的，我们不能让我们复合出来的channel有任何的元素。
		case 0:
			return nil
		case 1:  // 第二个结束条件是如果我们的可变切片只包含一个元素的时候，我们就返回那个元素。
			return channels[0]
		}

		orDone := make(chan interface{})
		// 这是函数的主体，以及递归发生的地方。我们启动了一个新的goroutine
		// 来让我们的channel 可以不被阻塞地等待消息的到来。
		go func() {
			defer close(orDone)

			switch len(channels) {
			// 因为我们所进行的循环迭代调用方式的原因，每次递归调用需要拥有至少两个channel。
			// 为了让我们所需要处理的 goroutine 的数目可被限制，我们在仅有两个 channel 调用
			// "or" 的时候设置了一个特殊情况。
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			// 在这里，我们从输入的多个channel 的第三个项目开始递归的生成 or-channel
			// 然后从 "select" 语句中选择一个合适的 channel。
			// 这将形成一个由现有slice 的剩余部分组成的树并且返回第一个信号量。
			// ！！！为了使在建立这个树的 goroutine 退出的时候在树下的 goroutine 也可以跟着退出，
			// 我们将这个 orDone channel 也传递到了调用中。
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}

	// 此功能只是创建一个channel，当后续时间中指定的时间结束时将关闭该channel
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	// 大致来追踪来自or函数的channel何时开始阻塞
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	// 打印发生读取事件所消耗的时间。
	fmt.Printf("done after %v", time.Since(start))
}

/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/30 21:32 8月
 **/
package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(
		done <-chan interface{},
		pulseInterval time.Duration,
	) (<-chan interface{}, <-chan time.Time) {
		// 我们建立了一个发送心跳的 channel 我们把这个返回给 doWork
		heartbeat := make(chan interface{})
		results := make(chan time.Time)
		go func() {
			defer close(heartbeat)
			defer close(results)

			// 我们设定心跳的间隔时间为我们接到的 pulseInterval。每隔一个
			// pulseInterval 的时长都会有一些消费者读取这个channel
			pulse := time.Tick(pulseInterval)
			// 这是另一个用来模拟心跳行为的channel。我们选择的持续时间大于 pulseInterval,
			// 这样我们就能看到从 goroutine 中发出的一些心跳。
			workGen := time.Tick(2 * pulseInterval)

			sendPulse := func() {
				select {
				case heartbeat <-struct{}{}:
				default: // 这里我们加入一个默认语句。我们必须时刻警惕这样一个事实：
				// 可能会没有人接收我们的心跳。从goroutine 发出的信息是重要的，但心跳却不一定重要。
				}
			}

			sendResult := func(r time.Time) {
				for {
					select {
					case <-done:
						return
					case <-pulse:  // 就像done channel 一样，当你执行发送或接收时，你也需要一个处理并
					// 发送心跳的 case 语句
						sendPulse()
					case results <- r:
						return
					}
				}
			}

			for {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case r := <-workGen:
					sendResult(r)
				}
			}
		}()
		return heartbeat, results
	}
	done := make(chan interface{})
	// 声明一个标准的 done channel 并在10s 后关闭。这给goroutine 留出了一些执行逻辑的时间。
	time.AfterFunc(10 * time.Second, func() {
		close(done)
	})
	// 这里我们设置了超时时间。我们使用此方法将心跳间隔与超时时间联系起来。
	const timeout = 2 * time.Second
	// 我们这里将心跳时间设为 timeout/2。这使我们的心跳有额外的响应时间，以便我们的超时不太敏感。
	heartbeat, results := doWork(done, timeout / 2)
	for {
		select {
		// 这里我们处理心跳。当没有消息时，我们至少知道每过 timeout / 2 的时间会从心跳 channel 发出一条消息。
		// 如果我们什么都没有收到，我们更知道是goroutine 本身出了问题。
		case _, ok := <-heartbeat:
			if ok == false {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results:
			if ok == false {
				return
			}
			fmt.Printf("results %v\n", r.Second())
		// 如果我们没有收到心跳或其他消息，就会超时。
		case <-time.After(timeout):
			return
		}
	}

}

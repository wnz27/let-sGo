/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/31 00:37 8月
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

			pulse := time.Tick(pulseInterval)
			workGen := time.Tick(2 * pulseInterval)
			sendPulse := func() {
				select {
				case heartbeat <-struct{}{}:
				default:
				}
			}

			sendResult := func(r time.Time) {
				for {
					select {
					case <-pulse:
						sendPulse()
					case results <- r:
						return
					}
				}
			}
			// 这是我们模拟的问题。所以它不是无限循环的，不需要我们动手停止
			// 就像墙面的例子一样，我们只会循环两次
			for i := 0; i < 2; i ++ {
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
	time.AfterFunc(10 * time.Second, func() {
		close(done)
	})
	const timeout = 2 * time.Second
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
			fmt.Println("worker goroutine is not healthy!")
			return
		}
	}
}

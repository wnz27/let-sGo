/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/31 01:02 8月
 **/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	doWork := func(done <-chan interface{}) (<-chan interface{}, <-chan int) {
		// 创建一个缓冲区大小为1 的heartbeat channel。这确保了即使没有及时接收发送的消息，
		// 至少也会发出一个心跳。
		heartbeatStream := make(chan interface{}, 1)
		workStream := make(chan int)
		go func() {
			defer close(heartbeatStream)
			defer close(workStream)

			for i:=0; i< 10; i++ {
				// 我们为心跳设置了一个单独的select 块。我们希望将发送results 和心跳分开，
				// 因为如果接收者没有准备好接收结果，作为替代它将接收到一个心跳，而代表当前结果的值将会丢失。
				// 由于我们有默认逻辑，所以这里也没有包含对 done channel 的处理。
				select {
				case heartbeatStream <- struct{}{}:
				// 为了防止没人接收我们的心跳，我们增加了默认逻辑。因为我们的 heartbeat channel
				// 创建时有一个容量的缓冲区，所以如果有人正在监听，但是没有及时收到第一个心跳，
				// 接收者仍然可以收到心跳。
				default:
				}
				select {
				case <-done:
					return
				case workStream <- rand.Intn(10):
				}
			}
		}()
		return heartbeatStream, workStream
	}

	done := make(chan interface{})
	defer close(done)

	heartbeat, results := doWork(done)
	for {
		select {
		case _, ok := <-heartbeat:
			if ok {
				fmt.Println("pulse")
			} else {
				return
			}
		case r, ok := <-results:
			if ok {
				fmt.Printf("results %v\n", r)
			} else {
				return
			}
		}
	}
}

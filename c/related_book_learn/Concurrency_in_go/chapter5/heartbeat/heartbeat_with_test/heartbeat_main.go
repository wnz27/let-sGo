/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/9/2 00:28 9月
 **/
package heartbeat_with

import (
	"time"
)

func DoWork(
	done <-chan interface{},
	nums ...int,
) (<-chan interface{}, <-chan int) {
	heartbeat := make(chan interface{}, 1)
	intStream := make(chan int)
	go func() {
		defer close(heartbeat)
		defer close(intStream)

		// 我们模拟在 goroutine 开始工作之前的某种延迟。在实践中，这可能是
		// 各种各样的问题，而且无法确定。我曾经见过CPU 负载过高、磁盘抢占、
		// 网络延迟和 goblins 造成的延迟
		time.Sleep(2 * time.Second)

		for _, n := range nums {
			select {
			case heartbeat <- struct{}{}:
			default:
			}
			select {
			case <- done:
				return
			case intStream <- n:
			}
		}
	}()

	return heartbeat, intStream
}


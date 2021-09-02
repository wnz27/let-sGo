/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/9/3 02:02 9月
 **/
package hb_with_test_iters

import (
	"testing"
	"time"
)

func DoWork(
	done <-chan interface{},
	pulseInterval time.Duration,
	nums ...int,
) (<-chan interface{}, <-chan int) {
	heartbeat := make(chan interface{}, 1)
	intStream := make(chan int)
	go func() {
		defer close(heartbeat)
		defer close(intStream)

		time.Sleep(2 * time.Second)

		pulse := time.Tick(pulseInterval)

		numLoop:  // 使用一个跳转标志来简化内部循环
		for _, n := range nums {
			// 我们需要两个循环, 一个循环遍历数列，内部循环持续执行，直到intStream 中的数字成功发送。
			for {
				select {
				case <-done:
					return
				case <-pulse:
					select {
					case heartbeat <- struct{}{}:
					default:
					}
				case intStream <- n:
					// 跳回 numLoop 标签继续执行外部循环
					continue numLoop
				}
			}
		}
	}()
	return heartbeat, intStream
}

func TestDoWork_GeneratesAllNumbers(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5}
	const timeout = 2 * time.Second

	heartbeat, results := DoWork(done, timeout / 2, intSlice...)

	<- heartbeat  // 等待第一次心跳到达，来确认 goroutine 已经进入了循环。

	i := 0
	for {
		select {
		case r, ok := <-results:
			if ok == false {
				return
			} else if expected := intSlice[i]; r != expected {
				t.Errorf(
					"index %v: expected %v, but received %v,",
					i,
					expected,
					r,
				)
			}
			i ++
		case <-heartbeat:  // 接收心跳(timeout / 2)，防止超时
		case <-time.After(timeout):
			t.Fatal("test timed out")
		}
	}
}


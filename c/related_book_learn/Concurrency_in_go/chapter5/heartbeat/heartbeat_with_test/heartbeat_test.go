/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/9/2 00:41 9月
 **/
package heartbeat_with

import (
	"testing"
	"time"
)

func TestDoWorkGeneratesAllNumbers(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5}
	_, results := DoWork(done, intSlice...)

	for i, expected := range intSlice {
		select {
		case r := <-results:
			if r != expected {
				t.Errorf(
					"index %v: expected %v, but received %v,",
					i,
					expected,
					r,
				)
			}
		// 在这里我们设置一个合理的超时时间，避免 goroutine 陷入死锁
		case <-time.After(1 * time.Second):
			t.Fatal("test time out")
		}
	}
}

func TestDoWork_GeneratesAllNumbers(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5}
	heartbeat, results := DoWork(done, intSlice...)

	<- heartbeat  // 这里我们等待 goroutine 开始处理迭代信号

	i := 0
	for r := range results {
		if expected := intSlice[i]; r != expected {
			t.Errorf("index %v: expected %v, but received %v,", i, expected, r)
		}
		i ++
	}
}

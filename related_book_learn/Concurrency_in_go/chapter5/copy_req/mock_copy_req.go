/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/9/3 23:52 9月
 **/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	doWork := func(
		done <-chan interface{},
		id int,
		wg *sync.WaitGroup,
		result chan<- int,
	) {
		started := time.Now()
		defer wg.Done()

		// 模拟随机负载
		simulatedLoadTime := time.Duration(1 + rand.Intn(5)) * time.Second
		select {
		case <- done:
		case <-time.After(simulatedLoadTime):
		}
		select {
		case <-done:
		case result <- id:
		}

		took := time.Since(started)
		// 显示处理程序需要多长时间
		if took < simulatedLoadTime {
			took = simulatedLoadTime
		}
		fmt.Printf("%v took %v\n", id, took)
	}

	done := make(chan interface{})
	result := make(chan int)

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i ++ {
		// 我们启动10个处理程序来处理请求。
		go doWork(done, i, &wg, result)
	}

	// 获得处理程序组的第一个返回值
	firstReturned := <-result
	// 我们取消其余的处理程序，以保证他们不会继续做多余的工作。
	close(done)
	wg.Wait()

	fmt.Printf("Received an answer from #%v\n", firstReturned)
}

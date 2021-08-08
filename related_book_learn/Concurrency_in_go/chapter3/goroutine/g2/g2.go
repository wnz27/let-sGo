/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/8 17:17 8月
 **/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
GC 并没有UI手被丢弃的goroutine 比如
go func () {
	// 将永远阻塞的操作
}
// 开始工作

这里的goroutine 将一直存在直到进程退出， 这个问题后面再讲。

现在利用这一点来测算goroutine 的大小

 */
func main() {
	// 我们将goroutine 不被gc 的事实与运行时的自省能力结合起来，并测算在goroutine创建之前和之后分配的内存数量：
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}
	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() {wg.Done(); <-c}

	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()

	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()

	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after - before)/numGoroutines/1000)
}

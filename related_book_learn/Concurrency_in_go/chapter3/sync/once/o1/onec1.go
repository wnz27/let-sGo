/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/9 01:43 8月
 **/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	increment := func() {
		count ++
	}
	var once sync.Once

	var incrementes sync.WaitGroup
	incrementes.Add(100)
	for i:=0; i < 100; i ++ {
		go func() {
			defer incrementes.Done()
			once.Do(increment)
		}()
	}
	incrementes.Wait()
	fmt.Printf("Count is %d \n", count)

	// sync.Once 是一种类型，它在内部使用一些sync 原语，以确保即使在不同的goroutine上，也只会调用一次
	// Do方法处理传递进来的函数。

	// 看看Go语言标准库它自己使用了多少次这个原语：
	// grep -ir sync.Once $(go env GOROOT)/src |wc -l
	// 128 次

}

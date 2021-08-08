/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/8 18:22 8月
 **/
package main

import (
	"fmt"
	"sync"
	"time"
)

/*
当你不关心并发操作的结果，或者你有其他方法来收集它们的结果时，WaitGroup是等待一组并发操作完成的好方法。

如果这个两个条件都不满足，那建议使用channel 和 select
 */
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait()
	fmt.Println("All goroutine complete.")

	/*
	输出如下
	2nd goroutine sleeping...
	1st goroutine sleeping...
	All goroutine complete.
	*/

	/*
	可以将WaitGroup视为一个并发-安全的计数器：调用通过传入的整数执行add方法增加计数器的增量，
	并调用Done方法对计数器进行递减。Wait阻塞，直到计数器为零。
	 */

	/* TODO 注意
	    添加调用 add 是在它们帮助跟踪的goroutine之外完成的。
		如果不这样，我们就会引入一种竞争条件，因为在本章前面"goroutine" 中，
		我们并不能保证goroutine何时被调度，可以在goroutine开始调度前调用Wait方法。
		如果将Add方法放在goroutine的闭包中，那么Wait调用可能会直接返回，而不会阻塞，
		因为add调用还未发生。
	 */

	/*
	通常情况下， 尽可能的向他们正在帮助追踪的goroutine中添加尽可能多的信息，
	但有时你会发现只调用一次Add来追踪一组goroutine。
	我通常在这样的循环之前执行这样的操作。
	 */
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}

	const numGoroutines = 5
	var wg2 sync.WaitGroup
	wg2.Add(numGoroutines)
	for i := 0; i < numGoroutines; i ++ {
		go hello(&wg2, i + 1)
	}
	wg2.Wait()
	/*
	输出可能如下：
	Hello from 5!
	Hello from 1!
	Hello from 2!
	Hello from 3!
	Hello from 4!
	*/
}

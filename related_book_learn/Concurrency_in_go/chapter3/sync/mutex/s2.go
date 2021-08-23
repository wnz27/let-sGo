/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/8 18:58 8月
 **/
package main

import (
	"fmt"
	"sync"
)

/*
Mutex 提供了一种安全的方式来表示对这些共享资源的独占访问。
为使用一个资源，channel通过通信共享内存，而Mutex通过开发人员的约定同步方位共享内存。
可以通过使用Mutex 对内存进行保护来协调退内存的访问。
 */
func main() {
	// 一个增加减少共享值的例子。
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count ++
		fmt.Printf("Increment: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count --
		fmt.Printf("Dcrement: %d\n", count)
	}

	// 增量
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i ++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	// 减量
	for i := 0; i <= 5; i ++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")

}

/**
 * @project let-sGo
 * @Author 27
 * @Description
 * @Date 2021/8/8 12:44 8月
 **/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	salutation := "hello"

	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()

	wg.Wait()

	fmt.Println(salutation)

	// goroutine 在它们所创建的相同的地址空间内执行，因此我们的程序打印出 "welcome" 这个词。

	// 让我们再看一个例子
	var wg2 sync.WaitGroup

	for _, salutation2 := range []string{"h1", "g1", "d1"} {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			fmt.Println(salutation2)
		}()
	}
	wg2.Wait()
	/*
	goroutine 正在运行一个闭包，该闭包使用变量salutation时，字符串迭代已经结束。
	当循环迭代时，salutation被分诶到的slice literal中下一个字符串值。
	因为计划冲的goroutine 可能在未来的任何时间点运行，它不确定在goroutine中会打印出什么值。
	在goroutine 开始之前循环有很高的的概率会退出。这意味着变量salutation 的值不在范围之内。
	这意味着变量salutation的值不在范围之内。然后会发生什么呢？
	goroutine还能引用一些已经超出范围的东西么？goroutine不会访问那些可能被垃圾回收的内存么？

	这是关于内存管理有趣的点，Go语言运行时会足够小心的将对变量salutation值的引用仍然保留，由内存转移到堆，
	以便goroutine 可以继续访问它。
	 */

	// 解决这个问题的方法是将循环的每一次，把salutation当时的副本传递到闭包中，这样当goroutine运行时，
	// 它将从循环的迭代中操作数据:
	var wg3 sync.WaitGroup
	for _, salutation3 := range []string{"aa", "bb", "cc"} {
		wg.Add(1)
		go func(s string) {  // 声明参数，把原来的变量显式的映射到闭包中。
			defer wg.Done()
			fmt.Println(s)
		}(salutation3) // 将当前迭代的变量传递给闭包。函数传参是copy副本
	}
	wg3.Wait()


}

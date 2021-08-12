/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/9 00:09 8月
 **/
package main

import (
	"fmt"
	"sync"
)

/*
为了解使用Broadcast的方法，让我们假设正在创建一个带有按钮的 GUI 应用程序。
我们想注册任意数量的函数，当该按钮被单机时，它将运行。

Cond 可以完美胜任，因为我们可以使用它的Broadcast方法通知所有注册的处理程序。
让我们看看它的例子：
 */

// 定义了一个Button类型，包含一个条件，Clicked
type Button struct {
	Clicked *sync.Cond
}

func main() {
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	// 定义了一个遍历构造函数，它允许我们注册函数处理来自条件的信号。
	// 每个处理程序都是在自己的goroutine上运行，订阅不会退出，
	// 直到 goroutine 被确认运行为止
	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	// 我们为鼠标按键时间设置了一个处理程序。
	// 它反过来调用Cond上的Broadcast，让所有的处理程序都知道鼠标按键已经被单机了（
	// 更健壮的实现将先检查它是否已经被抑制 todo 怎么做？？？）。
	var clickRegistered sync.WaitGroup
	// 创建一个waitGroup 这只是为了确保我们的程序在写入stdout之前不会退出。
	clickRegistered.Add(3)
	// 注册一个处理程序，当单机按键时，它将模拟最大化按钮的窗口。
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})
	// 模拟单机时显示对话框
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	// 模拟一个用户通过单机应用程序的按钮来单机鼠标按键
	button.Clicked.Broadcast()

	clickRegistered.Wait()

	/*
	可以看到，在Clicked Cond 上调用Broadcast，所有三个处理程序都将运行。
	如果不是clickRegistered 的 waitGroup, 我们可以调用button.Clicked.Broadcast() 多次
	并且每次都调用三个处理程序。
	这是channel不太容易做到的，因此是利用Cond类型的主要原因之一。
	 */

	/*
	与sync 包中所包含的大多数其他东西一样，Cond的使用最好被限制在一个紧凑的范围中，
	或者是通过封装它的类型来暴露在更大的范围内。
	 */

}

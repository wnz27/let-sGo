/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/9 01:52 8月
 **/
package main

import "sync"

func main() {
	// 看看这个例子会发生什么
	var onceA, onceB sync.Once
	var initB func ()
	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) } // 1

	onceA.Do(initA)  // 2
	// 1 这行调用在 2 返回之前不能进行
	/*
	这个程序将会死锁，因为在1调用的Do直到2调用Do并退出后才会继续，
	这是死锁的典型例子。

	对一些人来说这可能有些反直觉，因为它看起来好像我们使用的sync.Once 是为了防止
	多重初始化，但是sync.Once唯一能保证的是你的函数只被调用一次。有时
	这是通过死锁程序和暴露逻辑中的缺陷来完成的，在这个例子中是一个循环引用。
	 */
}

/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/8 19:54 8月
 **/
package main

import "sync"

/* cond
一个goroutine 的集合点，等待或者发布一个event
在这个定义中，一个event 是两个或者两个以上的goroutine之间的任意信号，除了它已经发生的event外，没有任何额外的信息。
通常情况下，在goroutine继续执行之前，你需要等待其中一个信号。
如果我们要研究如何在没有Cond类型的情况下实现这一目标，一个简单的办法就是使用无限循环。

for conditionTrue() == false {
}

然而这将消耗一个CPU 核心的所有周期，为了解决这个问题，我们可以引入一个time.Sleep

for conditionTrue() == false {
	time.Sleep
}

这样更好，但它仍然低效，而且你必须弄清楚要等待多久：太长，会人为的降低性能；
太短，会不必要等待消耗太多的CPU时间。

如果有一种方法可以让goroutine有效地等待，直到它发出信号并检查它的状态，那就好了。
这正是Cond类型为我们做的。
 */

func conditionTrue() bool {
	return true
}

func main() {
	// todo 这个例子没太理解
	//使用Cond，我们可以这样编写前面的例子：

	// 实例化一个新的cond， NewCond 函数创安一个类型，满足sync.Locker 接口。
	// 这使得cond类型能够以一种并发安全的方式与其他goroutine协调。
	c := sync.NewCond(&sync.Mutex{})

	// 锁定这个条件，这是必要的，因为在进入Locker的时候，执行wait会自动执行Unlock
	c.L.Lock()
	for conditionTrue() == false {
		// 等待通知，条件已经发生这是一个阻塞通信，main goroutine 将被暂停。
		c.Wait()
	}
	// 我们认为这个条件Locker执行解锁操作。这是必要的，因为当执行wait退出操作的时候，
	// 它会在Locker上调用Lock方法
	c.L.Unlock()
}
/*
这种方法效率更高。
注意，调用Wait 不只是阻塞，它挂起了当前的goroutine，允许其他goroutine在OS线程上运行。
当你调用Wait时，会发生一些其他事情：进入Wait后，在Cond变量的Locker上调用Unlock方法，
在退出Wait时，在Cond变量的Locker上执行Lock方法。在我看来，这需要慢慢习惯，它实际上是方法的
一个隐藏的副作用。担起来我们在等待条件发生的时候一直持有这个锁，但是事实并非如此。
当浏览代码时，需要留意这个模式。

 */

# 防止goroutine泄露

goroutine廉价且易于创建，这是让Go语言这么富有成效的原因之一。
运行时将多个goroutine复用到任意数量的操作系统线程，以便我们不必担心该抽象级别。

但是goroutine还是需要消耗资源，而且goroutine不会被运行时垃圾回收，所以无论
goroutine所占的用的内存有多么的少，我们都不希望我们的进程对此没有感知。

> 那么我们如何去确保他们被清理干净呢？
从头来思考这个问题：为什么一个goroutine需要存在呢？

在第二章，我们确定goroutine代表可能或不可能相互平行运转的工作单位。

goroutine有以下几种方式被终止：
- 当它完成了它的工作
- 因为不可恢复的错误，它不能继续工作。
- 当它被告知需要终止工作。

我们可以很简单的使用前两种方方法，因为这两种方法就隐含在你的算法中，
但是"取消工作"又是怎样工作的呢？由于网络影响，事实证明这是最重要的一点：
如果你开始了一个goroutine，最有可能以某种有组织的方式与其他几个goroutine合作。
我们甚至可以将这种相互连接表现为一个图表：
> 子goroutine是否应该继续执行可能是以许多其他goroutine状态的认知为基础的。

goroutine（通常是main goroutine）具有这种完整的语境知识应该能够告诉其子goroutine终止。
我们将在下一章继续研究大规模的goroutine的相互依赖关系，但现在让我们考虑如何确保一个子goroutine被清理。
让我们从一个简单的goroutine泄露开始：
```go
package main

import "fmt"

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				// 做些有趣的事
				fmt.Println(s)
			}
		}()
		return completed
	}

	doWork(nil)
	// 也许有其他的操作需要进行
	fmt.Println("Done")
}
```
[【demo】](g1_leak_eg/g1.go)

在这里我们看到 main goroutine 将一个空的channel传递给了doWork。
因此，string channel 永远也不会获得任何string，并且包含doWork函数的goroutine
会一直在程序的生命周期内保持在内存中
（我们甚至可能会在将doWork内的goroutine与main goroutine 进行结合的时候造成死锁）

在这个例子中，这个过程的生命周期是十分短暂的，但是在真正的程序中，
goroutine应该很常见的会在一个长寿的程序初始化的时候就被启动。

最糟糕的情况下，main goroutine 可能会在其生命周期内持续的将其他的 goroutine 设置为
自旋（TODO 这个翻译没懂），这会导致内存利用率的下降。

将将父子goroutine进行成功整合的一种方法就是在父子goroutine之间建立一个"信号通道"，
让父goroutine可以向子goroutine发出取消信号。
按照惯例，这个信号通道是一个名为done的只读channel。父goroutine将该channel传给
子goroutine，然后在想要取消子goroutine时关闭该channel。例如
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface {} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					// 做一些有意思的事情
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		// 在 1s 之后取消本操作
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-terminated
	fmt.Println("Done")
}
```
输出：
```shell
Canceling doWork goroutine...
doWork exited.
Done
```
[【demo】](g2_done_eg/g2.go)

可以看到，尽管我们给我们的字符串channel中传递了nil，我们的goroutine仍然成功退出。
与之前的例子不同，在这个例子中，我们加入了两个goroutine，但是没有造成死锁。
这是因为在我们加入两个goroutine之前，我们创建了第三个goroutine来在doWork 执行
1s之后取消doWork中的goroutine。我们成功消除了我们的goroutine泄露！




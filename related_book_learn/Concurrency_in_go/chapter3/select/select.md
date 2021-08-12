# select 语句
select语句是将channel绑定在一起的粘合剂，这就是我们如何在一个程序中组合channel以形成更大的
抽象事务的方式。

如果channel是将goroutine连接在一起的黏合剂，那么声明select语句是做什么的呢？

声明select语句是一个具有并发性的Go语言程序中最重要的事情之一，这并不是夸大其词。

在一个系统中两个或多个组件的交集中，可以在本地、单个函数
或类型以及全局范围内找到select语句绑定在一起的channel。

除了连接组件之外，在程序中的这些关键节点上，select语句可以帮助安全地将channel与
诸如取消、超时、等待和默认值之类的概念结合在一起。
相反，如果select语句是程序的通用语言，它们只处理channel，那么程序的组件应该如何协调？
我们将在第五章专门研究这个问题（提示：更推荐使用channel）

那么这些强大的select语句是什么呢？我们如何使用它们，它们是如何工作的？
让我们先把它放出来。这里有一个简单的例子：
```go
package main

func main() {
	var c1, c2 <- chan interface{}
	var c3 chan<- interface{}

	select {
	case <- c1:
		//执行某些逻辑
	case <- c2:
		// 执行某些逻辑
	case c3<- struct {}{}:
		// 执行某些逻辑
	}
}
```
像一个选择模块，一个select模块包含系列的case语句，这些语句可以保护一系列语句。
然而，这就是相似之处。
与switch块不同，select块中的case语句没有测试顺序，如果没有满足任何条件，执行也不会失败。

相反，所有的channel读取和写入都需要查看是否有任何一个已准备就绪可以用的数据（实际运行情况要更复杂一些，我们会在第6章提及）：
在读取的情况下关闭channel，以及写入不具备下游消费能力的channel。
如果所有channel都没有准备好，则执行整个select语句模块。当一个channel准备好了，这个操作就会继续，
它相应的语句就会执行。

看一个简单示例：
```go
func main() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5*time.Second)  // 等待5shou关闭channel
		close(c)
	}()
	
	fmt.Println("Bocking on read ...")
	select {
	case <-c:
		fmt.Printf("Unblock %v later.\n", time.Since(start))
	}
}
```
[【demo】](s1/s1.go)
输出：
```shell
Bocking on read ...
Unblock 5.002684552s later.
```
如你所见，在进入select模块后大约5s，我们就会解锁。
这是一种简单而有效的方法来阻止我们等待某事的发生，但如果我们思考一下，我们可以提出一些问题：
- **当多个channel有数据可供给下游读取的时候会发生什么？**
- **如果没有任何可用的channel怎么办？**
- **如果我们想要做一些事情，但是没有可用的channel怎么办？**

### 多个channel同时可用的这个问题似乎很有趣，试一下：
```go
package main

import "fmt"

func main() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int
	for i := 1000; i >= 0; i -- {
		select {
		case <-c1:
			c1Count ++
		case <-c2:
			c2Count ++
		}
	}
	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
}
```
[【demo】](s2/s2.go)
输出：
```shell
c1Count: 495
c2Count: 506
```
如你所见，在一千次迭代中，大约有一半时间从c1读取select语句，大约一半时间从c2读取。
这看起来很有趣，也许有点太巧了。 事实如此。

Go语言运行时将在一组case语句中执行伪随机选择。这就意味，在你的case语句集合中，
每一个都有一个被执行的机会。

（怎么做？）
乍一看这似乎并不重要，但背后的原因却非常有趣。让我们先做一个很明显得阐述：
Go语言运行时无法解析select语句的意图，也就是说，它不能推断出问题空间，或者说为什么
将一组channel组合到一个select语句中。

正因为如此，运行时所能做得最好的事情就是在平均的情况下运行良好。
一种很好的方法是将一个随机变量引入到等式中？？（在这种情况下，select后续的channel）
通过加权平均每个channel被使用的机会，所有使用select语句的程序将在平均情况下表现良好。

### 关于第二个问题：如果没有任何channel可用，会发生什么？
如果所有的channel都被阻塞了，如果没有可用的，但你可能不希望永远阻塞（会死锁的），可能需要超时机制。

Go语言的time包提供了一种优雅的方式，可以在select语句中很好的使用channel。这里有一个例子：
```go

```







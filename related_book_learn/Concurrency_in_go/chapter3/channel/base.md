CSP: Communication Sequential processes 通信顺序进程

Go的一个座右铭：使用通信来共享内存，而不是通过共享内存来通信。

## 基础

Channel就是由CSP派生的同步原语。

虽然他们可以用来同步内存访问，但是它们最好用于在goroutine之间传递信息。 任何大小程序中，channel 都非常有用，因为它们可以组合在一起。

就像河流一样，一个channel充当着信息传递的管道，值可以沿着channel传递，然后在下游读出。 出于这个原因，通常用"Stream"来做chan变量名的后缀。

当你使用channel时，你会将一个值传递给一个chan变量，然后你程序中的某个地方将它从channel中读出。

程序中不同部分不需要相互了解，只需要在channel所在的内存中引用相同的位置即可。 这可以通过对程序上下游的channel的引用来完成。

创建一个channel非常简单。可以使用 := 操作就可以在一个语句中创建channel 但因为你需要经常声明channel，因此将这两步操作拆分为单个语句是很有用的：

```go
var dataStream chan interface{}  // 1
dataStream = make(chan interface{}) // 2
```

- 1、声明一个channel，因为我们声明的类型是空接口，所以说它的类型是interface{}
- 2、使用内置的make函数实例化channel

> 这个例子定义了一个channel（dataStream），任何值都可以写入或者读取（因为我们使用了空接口）

## 单向

channel也可以声明只支持单向的数据流，也就是说，可以定义一个channel只支持发送或者接收信息。

要声明一个单向channel，只需要包括 <- 操作符。

要声明和实例化一个只能读取的channel，将 <- 操作符放在左侧，就像这样:

```go
var dataStream <-chan interface{}
dataStream = make(<-chan interface{})
```

要声明并创建一个只能发送的channel，将 <- 操作符放在右侧，就像这样:

```go
var dataStream chan<- interface{}
dataStream = make(chan<- interface{})
```

通常不会看到单向的channel 实例化，但是会经常看到它们的用作函数参数和返回类型， 这是非常有用的。 因为当需要时，Go语言会隐式的将双向channel转换为单向channel。这里有一个例子：

```go
var receiveChan <-chan interface{}
var sendChan chan<- interface{}
dataStream := make(chan interface{})
// 有效语法：
receiveChan = dataStream
sendChan = dataStream
```

我们既然可以创建interface{} 类型的chan，意味着我们也可以创建一个更严格的类型来约束它可以传递的 数据类型。 这是一个整数的例子：

```go
intStream := make(chan int)
```

为了使用channel， 我们将再次使用<-操作符。 通过将<-操作符放到channel的【右边】实现【发送】操作， 通过将<-操作符放到channel的【左边】实现【接收】操作。

另一种思考方式是 数据流向箭头所指的方向的变量。

```go
stringStream := make(chan string)
go func () {
stringStream <- "Hello channels!" // 我们将字符串文本传递到 stringStream channel
}()
fmt.Println(<-stringStream) // 我们读取channel的字符串字面量并将其打印到 stdout
```

你可以将数据传递给channel变量，并读取它的数据。

但是尝试将一个值写入只读的channel是错误的， 并且从只可以写的channel读取值也是错误的。 如果我们这样写了，编译器会直接告诉我们非法：

```go
writeStream := make(chan<- interface{})
readStream := make(<-chan interface{})

<-writeStream
readStream <- struct{}
```

编译不能通过：

```shell
invalid operation: <-writeStream(receive from send-only type chan<- interface{})
invalid operation: readStream <- struct {} literal (send to receive-only type <-chan interface{})
```

这是Go语言的类型系统的一部分，它允许我们在处理并发原语时使用type-safety。正如我们稍后将在本节中看到的，这是一种强大的方法， 可以声明我们的API并构建可组合的、易于推理的逻辑程序。

## goroutine 退出问题

之前抛出过一个问题，就是goroutine 是被动调度的，没有办法保证它会在程序退出之前运行。

但是前面的示例（如下）是完整的，没有省略任何代码，这个示例为什么匿名的goroutine在main goroutine之前就完成运行了呢？

```go
stringStream := make(chan string)
go func () {
stringStream <- "Hello channels!" // 我们将字符串文本传递到 stringStream channel
}()
fmt.Println(<-stringStream) // 我们读取channel的字符串字面量并将其打印到 stdout
```

因为Go语言中channel是阻塞的。这意味着只要channel内的数据被消费后，新的数据才能写入， 而任何试图从空channel读取数据的goroutine将等待至少一条数据被写入channel后才能读到。
上面例子正好符合这个要求。写入成功之前goroutine不会退出，main goroutine也不会瞬间执行完。

如果不正确的构造程序，会导致死锁。看下面示例：

```go
package main

import "fmt"

func main() {
	stringStream := make(chan string)
	go func() {
		if 0 != 1 {
			return
		}
		// 因为上面的条件，下面不会写入成功的
		stringStream <- "hello"
	}()
	fmt.Println(<-stringStream)
}
```

上面例子会直接panic:

```shell
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        /Users/fzk27/fzk27/let-sGo/related_book_learn/Concurrency_in_go/chapter3/channel/base.go:39 +0x156
Process finished with exit code 2
```

main goroutin在等一个值被写入stringStream channel, 但由于我们的逻辑，这将永远不会发生。 当匿名的goroutine退出时，检测到所有的goroutine都没有运行，并报了一个死锁。

本章后面会讲如何构造程序才能做到简单的防止这种死锁，在下一章中会讲如何完全避免这些问题。

[上面涉及部分demo](base.go)

## channel 消费问题

通过 <- 操作符的接受刑事也可选择返回两个值：

```go
package main

import "fmt"

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "hello"
	}()
	salutation, ok := <-stringStream
	fmt.Printf("(ok: %v): %v", ok, salutation)
	// 输出如下：
	// (ok: true): hello
```

第二个返回值是读取操作的一种方式，用于表示该channel上有新数据写入，或者是由closed channel生成的默认值。 closed channel是什么?

### 关闭channel

能够提示channel中是否会有新的值写入是非常有用的。 这有助于下游的程序知道什么时候消费、退出、给新的channel重新建立连接等。 我们可以给每个这样的的类型一个特殊标记，但是这将覆盖大多数开发人员的代码，
channel只是一个简单的数据传输channel，而不是数据类型的函数，所以关闭channel是一个比较 普通的操作，就好比哨兵说：嘿，上有不会写入任何有价值的数据了，想干嘛干嘛吧。 我们使用close关键字关闭一个channel

```go
valueStream := make(chan interface{})
close(valueStream)
```

有趣的是，我们也可以从一个已经关闭的channel读取数据。

```go
intStream := make(chan int)
close(intStream)
integer, ok := <- intStream
fmt.Printf("(%v): %v", ok, integer)
```

输出：(ok? false): 0  [demo](c1/c1.go)

注意我们从来没有把任何数据推送到channel上，立即关闭它。 我们仍然能够执行读取操作，事实上，尽管channel已经关闭，我们仍然可以继续在这个channel上执行读取操作。

这是为了支持一个channel有单个上游写入，有多个下游读取。 第二个返回值（也就是ok的值）是false，这表示我们收到的值是 int 或者 0, 而不是推到stream上的值。

这为我们提供了一些新的模式。第一个是从channel中获取。通过range作为参数遍历（for）并且在channel关闭时 自动中断循环。这允许对channel上的值进行简洁的迭代。让我们看一个例子:

```go
package main

import "fmt"

func main() {
	intStream := make(chan int)
	go func() {
		defer close(intStream) // 确保goroutine 退出之前channel是关闭的。这是一个常见的模式。
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()
	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}
}
```

输出: 1 2 3 4 5   [demo](c2/c2.go)

该循环不需要退出条件，并且range方法不返回第二个bool值。 处理一个已关闭的channel的细节可以让你保持循环简洁。

关闭channel也是一种同时给多个goroutine发信号的方法。 如果有n个goroutine在一个channel上等待， 而不是在channel上写n次来打开每个goroutine，你可以简单的关闭channel。
由于一个被关闭的channel可以被无数次读取，所以不管有多少goroutine在等待它， 关闭channel都比执行n次更合适，也更快。

这里有一个例子， 可以同时打开多个goroutine:

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin // goroutine会一直等待，直到它被告知可以继续
			fmt.Printf("%v has begun \n", i)
		}(i)
	}
	fmt.Println("Unblocking goroutines ......")
	close(begin) // 关闭channel 从而同时打开所有的goroutine
	wg.Wait()
}
```
可以看到我们关闭begin 的channel之前，所有的goroutine都没有开始执行，
```shell
Unblocking goroutines ......
1 has begun 
4 has begun 
3 has begun 
2 has begun 
0 has begun
```
[demo](c3/c3.go)

使用之前的sync.Cond的Broadcast也能实现。

但是正如我们讨论过的，channel是可以组合的，
这是作者最喜欢的一种在同一时间打开多个goroutine的方法
（个人思考：对比Broadcast看起来会更简单，可读性更好）

## buffered channel

它是在实例化时提供容量的channel。这意味着即使没有在channel执行读取操作，goroutine仍然
可以执行n写入，其中n是缓冲channel的容量。简单示例如下：
```go
var dataStream chan interface{}
dataStream = make(chan interface{}, 4)
// 创建一个有4容量的缓冲channel。
//这意味着我们可以把4条数据放到channel上，不管它是否被读取
```

再一次，我将实例化分解成两行，这样就可以看到一个缓冲的channel的声明与一个没有缓冲的channel
没有什么不同。
这有点意思，因为它意味着goroutine可以控制实例化一个channel时是否需要缓冲。

这表明，创建一个channel应该与goroutine紧密耦合，而goroutine将会在它上面执行写操作，
这样我们就可以更容易地推断它的行为和性能。稍后讨论这个问题。

没有缓冲的channel也可被定义为缓冲channel，只是缓冲为0，如下两个定义等效：
```go
a := make(chan int)
b := make(chan int, 0)
```
这俩channel都是具有0容量的int channel。

请记住当我们讨论阻塞时，【如果说channel是满的，那么写入阻塞】
【如果channel是空的，则读取是阻塞的】你能读什么呢？

Full 和 empty 是容量或换乘区大小的函数。

无缓冲channel的容量为0，因此在任何写入之前channel已经满了。

一个没有下游接受的容量为4的缓冲channel在写入4次之后就满了，并且在写入第5次的时候阻塞，
因为它没有其他地方放置第五个元素。与无缓冲的channel一样，缓冲channel仍然阻塞，
channel为空或满的前提条件是不同的。

通过这种方式，缓冲channel是一个内存中的FIFO队列，用于并发进程进行通信。

> 如果一个缓冲channel是空的，并且有一个下游接收，那么缓冲区将被忽略，
并且该值将直接从发送方传递到接收方，在实践中这是透明的，但是对了解缓冲channel的配置是值得的

缓冲channel在某些情况下是有用的，但是应该小心地创建它们。在下一章中，我们将看到，
缓冲channel很容易成为一个不成熟的优化，并且使隐藏的死锁更不容易发生（or发现？）
这听起来像个好事，但是我宁愿第一次写代码的时候发现死锁，而不是在生产系统崩溃的时候才发现。

来看一个完整的示例：
```go
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var stdoutBuff bytes.Buffer  // 创建一个内存缓冲区，以帮助减少输出的不确定性。
	// 它没有给我们任何保证，但它比直接写stdout 要快一些
	
	defer stdoutBuff.WriteTo(os.Stdout)  // 确保程序退出之前缓冲区内容需要被写入到stdout

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i:=0; i < 5; i ++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v. \n", integer)
	}
}
```
[demo](c4/c4.go)

本例中输出到stdout的顺序是不确定的。看下输出
```shell
Sending: 0
Sending: 1
Sending: 2
Sending: 3
Sending: 4
Producer Done.
Received 0. 
Received 1. 
Received 2. 
Received 3. 
Received 4. 
```
大概会发现我们的匿名goroutine是如何把它的5个结果都写到intStream上的，然后在
main goroutine将每个结果都推送出去之前就退出。

这是一个适合某些条件的优化例子：
如果goroutine写入一个channel的时候会知道它需要写多少条数据，
它可以创建一个容量是写数量的缓冲channel，然后尽快往channel里写数据。
当然这样做有一些注意事项，下章讨论。

### nil channel
channel的默认值是nil。



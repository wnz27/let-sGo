CSP: Communication Sequential processes 通信顺序进程

Go的一个座右铭：使用通信来共享内存，而不是通过共享内存来通信。

## 基础
Channel就是由CSP派生的同步原语。

虽然他们可以用来同步内存访问，但是它们最好用于在goroutine之间传递信息。
任何大小程序中，channel 都非常有用，因为它们可以组合在一起。

就像河流一样，一个channel充当着信息传递的管道，值可以沿着channel传递，然后在下游读出。
出于这个原因，通常用"Stream"来做chan变量名的后缀。

当你使用channel时，你会将一个值传递给一个chan变量，然后你程序中的某个地方将它从channel中读出。

程序中不同部分不需要相互了解，只需要在channel所在的内存中引用相同的位置即可。
这可以通过对程序上下游的channel的引用来完成。

创建一个channel非常简单。可以使用 := 操作就可以在一个语句中创建channel
但因为你需要经常声明channel，因此将这两步操作拆分为单个语句是很有用的：
```go
var dataStream chan interface{}  // 1
dataStream = make(chan interface{})  // 2
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

通常不会看到单向的channel 实例化，但是会经常看到它们的用作函数参数和返回类型，
这是非常有用的。
因为当需要时，Go语言会隐式的将双向channel转换为单向channel。这里有一个例子：
```go
var receiveChan <-chan interface{}
var sendChan chan<- interface{}
dataStream := make(chan interface{})
// 有效语法：
receiveChan = dataStream
sendChan = dataStream
```

我们既然可以创建interface{} 类型的chan，意味着我们也可以创建一个更严格的类型来约束它可以传递的
数据类型。
这是一个整数的例子：
```go
intStream := make(chan int)
```
为了使用channel， 我们将再次使用<-操作符。
通过将<-操作符放到channel的【右边】实现【发送】操作，
通过将<-操作符放到channel的【左边】实现【接收】操作。

另一种思考方式是 数据流向箭头所指的方向的变量。
```go
stringStream := make(chan string)
go func() {
	stringStream <- "Hello channels!"  // 我们将字符串文本传递到 stringStream channel
}()
fmt.Println(<-stringStream)  // 我们读取channel的字符串字面量并将其打印到 stdout
```
你可以将数据传递给channel变量，并读取它的数据。

但是尝试将一个值写入只读的channel是错误的，
并且从只可以写的channel读取值也是错误的。
如果我们这样写了，编译器会直接告诉我们非法：
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
这是Go语言的类型系统的一部分，它允许我们在处理并发原语时使用type-safety。正如我们稍后将在本节中看到的，这是一种强大的方法，
可以声明我们的API并构建可组合的、易于推理的逻辑程序。

## goroutine 退出问题
之前抛出过一个问题，就是goroutine 是被动调度的，没有办法保证它会在程序退出之前运行。

但是前面的示例（如下）是完整的，没有省略任何代码，这个示例为什么匿名的goroutine在main goroutine之前就完成运行了呢？
```go
stringStream := make(chan string)
go func() {
	stringStream <- "Hello channels!"  // 我们将字符串文本传递到 stringStream channel
}()
fmt.Println(<-stringStream)  // 我们读取channel的字符串字面量并将其打印到 stdout
```
因为Go语言中channel是阻塞的。这意味着只要channel内的数据被消费后，新的数据才能写入，
而任何试图从空channel读取数据的goroutine将等待至少一条数据被写入channel后才能读到。
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
main goroutin在等一个值被写入stringStream channel, 但由于我们的逻辑，这将永远不会发生。
当匿名的goroutine退出时，检测到所有的goroutine都没有运行，并报了一个死锁。

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
	salutation, ok := <- stringStream
	fmt.Printf("(ok: %v): %v", ok, salutation)
	// 输出如下：
	// (ok: true): hello
```
第二个返回值是读取操作的一种方式，用于表示该channel上有新数据写入，或者是由closed channel生成的默认值。
closed channel是什么?

### 关闭channel


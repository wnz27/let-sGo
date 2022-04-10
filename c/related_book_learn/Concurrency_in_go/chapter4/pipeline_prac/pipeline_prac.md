# 构建pipeline的最佳实践

channel 非常适合在Go语言中构建pipeline，因为它们满足了我们所有的基本要求。
- 它们可以接收并返回值
- 可以安全的被并行使用
- 它们可以被 range 语句所遍历
- 并且都被编程语句所 "具体化" (个人理解是显式声明)

让我们花点时间转换一下前面的例子来改用channel：
```go
package main

import "fmt"

func main() {
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}

	multiply := func(
	done <-chan interface{},
	intStream <-chan int,
	multiplier int,
	) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- i * multiplier:
				}
			}
		}()
		return multipliedStream
	}

	add := func(
	done <-chan interface{},
	intStream <-chan int,
	additive int,
	) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case addedStream <- i + additive:
				}
			}
		}()
		return addedStream
	}

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}
}
```
[【demo】](pipeline_with_channel/pipeline_with_channel.go)
输出：
```shell
6
10
14
18
```
它看起来已经输出了我们所期待的结果，但是代价是用了更多的代码。
我们究竟得到了什么？首先看看我们写的是什么？我们现在有三个函数而不是两个。
它们都在自己的函数体中启动了一个goroutine，并使用了我们在本章前面 
"防止 goroutine 泄露" 中建立的模式，通过一个 channel 表示该 goroutine 应该退出。
它们看起来都像是返回 channel，其中一些 channel 看起来像是包涵在另外一个 channel 中。
很有趣， 让我们开始进一步分解：
```go
done := make(chan interface{})
defer close(done)
```
我们的程序做的第一件事是创建一个done channel，并在 defer 语句中关闭它。
正如前面所讨论的那样，这可以确保我们的程序干净地离开，不会泄露goroutine。
没啥新东西。接下来，我们来看看函数generator：
```go
generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}
// ......
intStream := generator(done, 1, 2, 3, 4)
```
generator 函数接受一个可变的整数切片，构造一个缓存长度等于输入整数片段的整数 channel，
启动一个goroutine 并返回构造的channel。然后，在创建的goroutine 上，generator 函数
使用range 语句遍历传入的可变切片，并在其创建的channel 上发送切片的值。

需要注意的是，在这个channel 上进行数据发送与done channel 共享同一个 select 语句的。
再一次，这是我们在本章前面 "防止 goroutine泄露" 中建立的模式，以防止泄露goroutine。

简单来讲，生成器函数将一组离散值转换为一个channel 上的数据流。
也就是说，这种类型的函数被称为生成器。

当你在使用pipeline 的时候你将经常和上述的生成器函数打交道，因为在一个pipeline 的开始时，
你总是会有一堆需要被转化为 channel 的数据。我们将稍微介绍一些有趣的生成器的例子，
但我们先来完成对这个程序的分析，接下来，构建我们的pipeline：
```go
pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
```
这和我们一直在处理的pipeline是一样的：对于一串数字，我们将它们乘以2，加1，然后再乘以2。
这个channel 与我们之前的例子中实现函数公的pipeline 非常相似，但是在很重要的方式上是不同的。

首先我们正在使用pipeline。这是显而易见的，因为它允许两件事：在我们的pipeline的末尾，
我们可以使用范围语句（range）来提取值，并且在每个stage 我们可以安全地同时执行，因为我们的
输入和输出在并发上下文中是安全的。

这给我们带来了第二个不同之处：pipeline的每个stage 都在执行控制。
**这意味着任何stage 只需要等待其输入，并且能够发送其输出。** 【事实证明，这会产生巨大的影响，
我们将在本章后面"扇出，扇入"中进行介绍，但现在可以简单地注意到它允许我们的stage 相互独立地
执行某个片段时间。】

最后在我们的例子中，我们对这个 pipeline 进行了排序，并且通过系统获取了值：
```go
for v := range pipeline {
    fmt.Println(v)
}
```
（这里省略表格图片）

在允许完成之前，关闭done channel 会怎么样？这是由每个stage 中的如下两条来保证可能性的：
- 对传入的 channel 进行迭代，当传入的channel 已经 关闭了，range 语句也将会退出。
- 发送 channel 与 done channel 存在于同一个 select 语句中。

无论 pipeline stage 所处的 stage 的状态 如何（在等待传入的channel 还是等待发送完成），
关闭 done channel 都将会迫使整个pipeline stage 被终止。

这里有一个复发关系（没懂字面意思）。在pipeline 开始时，我们已经确定我们必须将离散值转换为
channel。在这个过程中有两点必须是可抢占的：
- 创建几乎不是瞬时的离散值（todo 可能翻译有问题，待考证）
- 在离散值的 channel 上进行发送。

你可以随意选择上述中的任意方案。在我们的例子中，在生成器函数中，离散值是通过遍历可变切片生成的，
他足够瞬时，不需要被抢占。
第二个是通过我们的 select 语句和 done channel 处理的，它确保发生器即使被阻塞试图写入
intStream 也是可抢占的。

在pipeline 的另一侧，通过引入（range）来保证最后的 stage 的可抢占性。
可抢占性是通过我们正在使用 range 语句进行遍历的channel将会在被强占时关闭，因此，
我们的 range 语句会在发生被强占的情况时退出。最终的stage 也是可被抢占的，因为我们
依赖的流本身是可被抢占的。

在pipeline 开始和结束之间，代码总是在使用 range 语句中遍历 channel，并且在一个包含
一个 done channel 的 select 语句中向其他的channel 发送消息。

如果一个stage 在从输入 channel 中获取值时被阻塞，该stage 将会在输入 channel 被关闭时解锁。

通过将channel 进行归纳，我们可以知道我们的channel 可能会因为它是我们内嵌的一个channel
或者是我们已经建立好可抢占的 pipeline 的开始而被关闭。当某个 stage 在发送数据时被阻塞，
多亏了 select 语句，它依旧是可被抢占的。

因此我们的整个 pipeline 始终可以通过关闭 done channel 来进行抢占。

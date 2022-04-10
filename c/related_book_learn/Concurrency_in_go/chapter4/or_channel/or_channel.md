# or-channel
有时候你想要将多个done channel合并成一个单一的done channel 来在这个复合 done channel
中的任意一个channel关闭的时候，关闭整个done channel。
编写一个执行这种耦合的选择语句是完全可以接受的，尽管很冗长。

但是有时候你无法知道你在运行时使用的done channel的数量。
在这种情况下，或者如果你只喜欢单线程，你可以使用or-channel模式将这些channel组合在一起。

这种模式通过递归和goroutine创建一个复合done channel。来看一下：
```go
package main

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	// 这就是我们的函数，它将一堆包含可变切片的channel 整合成一个单一的channel
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		// 因为这是一个循环函数，我们必须设置结束条件。首先，如果可变切片是空的，我们只返回一个空 channel。
		// 这里与我们传递进一个空channel 是相通的，我们不能让我们复合出来的channel有任何的元素。
		case 0:
			return nil
		case 1:  // 第二个结束条件是如果我们的可变切片只包含一个元素的时候，我们就返回那个元素。
			return channels[0]
		}

		orDone := make(chan interface{})
		// 这是函数的主体，以及递归发生的地方。我们启动了一个新的goroutine
		// 来让我们的channel 可以不被阻塞地等待消息的到来。
		go func() {
			defer close(orDone)

			switch len(channels) {
			// 因为我们所进行的循环迭代调用方式的原因，每次递归调用需要拥有至少两个channel。
			// 为了让我们所需要处理的 goroutine 的数目可被限制，我们在仅有两个 channel 调用
			// "or" 的时候设置了一个特殊情况。
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			// 在这里，我们从输入的多个channel 的第三个项目开始递归的生成 or-channel
			// 然后从 "select" 语句中选择一个合适的 channel。
			// 这将形成一个由现有slice 的剩余部分组成的树并且返回第一个信号量。
			// ！！！为了使在建立这个树的 goroutine 退出的时候在树下的 goroutine 也可以跟着退出，
			// 我们将这个 orDone channel 也传递到了调用中。
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}
}
```
这是一个相当简洁的函数，使你可以将任意数量的channel组合到单个channel中，
只要任何组件 channel 关闭或者写入，该channel就会关闭。
我们来看看如何使用这个功能。下面看一个简短的例子，它将经过一段时间后关闭的channel，
并将这些 channel 合并到一个关闭的单个channel中：
```go
// 此功能只是创建一个channel，当后续时间中指定的时间结束时将关闭该channel
sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
    // 大致来追踪来自or函数的channel何时开始阻塞
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
    // 打印发生读取事件所消耗的时间。
	fmt.Printf("done after %v", time.Since(start))
```
[【demo】](oc1/oc1.go)

输出：
```shell
done after 1.001544986s
```
需要注意的是，尽管我们在调用"or"的时候传入了多个拥有不同的关闭时间的channel，
但是，随着我们"1s 之后关闭的channel" 被关闭，整个由"or"函数所生成的channel都关闭了。
这是因为，无论它位于构建函数还是树中，它都会被首先关闭来确保依赖于它"关闭"的channel也会被关闭。

我们以使用更多的goroutine为代价，实现了这个简洁性。
f(x) = x/2，其中x是goroutine的数量，但你要记住Go语言的一个优点是能够快速创建，
调度和运行goroutine，并且该语言积极鼓励使用goroutine来正确建模问题。
担心在这里创建的goroutine的数量可能是一个不成熟的优化。
此外，如果在编译时你不知道你正在使用多少个 done channel，则将会没有其他方式可以
合并 done channel。

这种模式在你的系统中的【**模块交汇处非常有用**】。在这些交汇处，你的调用堆中
应该有复数种的用来取消goroutine 的决策树。
使用or 函数，你可以简单地将它们组合在一起并将其传递给堆栈。我们将在后面的"context 包"
中看到另一种做法，这也很好，也许更具描述性。

我们还将了解如何使用此模式的变体在 第五章 "复制请求" 中形成更复杂的模式。


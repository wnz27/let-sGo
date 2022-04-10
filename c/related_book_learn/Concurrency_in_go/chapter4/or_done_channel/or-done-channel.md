# or-done-channel
有时候，你需要处理来系统各个分散部分的channel。
与操作pipeline 所不同的是，你不能对你通过 done channel 
进行取消了的 channel 上的代码将会表现成什么样子做断言。
也就是说，你不知道你的goroutine 是否被取消，这意味着你正在读取的channel 将被取消。([从关闭的channel读的值一般是不期望的默认值](../../chapter3/channel/base.md))
处于这个原因，正如在本章前面 "防止 goroutine泄露" 中所阐述的那样，我们需要用channel 
中的select 语句来包装我们的读操作，并从已完成的channel 中进行选择。
这非常好，但是这样做需要的代码很容易读取，如下所示：
```go
for val := range myChan {
	// 用 val 执行某些操作
}
```
然后使用如下代码使其暴露出来：
```go
loop:
for {
	select {
	case <-done:
		break loop
    case maybeVal, ok := <-myChan:
    	if ok == false {
    		return // 或许从 for 循环中退出
        }
        // 用 maybeVal 执行某些操作
    }
}
```
这可能会很快就繁忙起来，特别是如果你有嵌套循环。继续使用 goroutine 编写更清晰的并发代码，
而不是过早优化，我们可以用一个 goroutine 来解决这个问题。
我们封装了详细信息，以便于其他的函数:
```go
package main

func main() {
	orDone := func(
		done, 
		c <-chan interface{},
	) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false { // 如果c关闭了
						return
					}
					// 不关闭则在起一个复用
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}
}
```
[【demo】](or_done.go)

这样做可以让我们回到简单的循环, 像这样：
```go
for val := range orDone(done, myChan) {
	// 用 val 执行某些操作
}
```
你可能会在你的代码中发现需要使用一系列select 语句的紧密循环的边界案例，
但我会鼓励你先尝试编写具有可读性的代码，并避免过早优化。



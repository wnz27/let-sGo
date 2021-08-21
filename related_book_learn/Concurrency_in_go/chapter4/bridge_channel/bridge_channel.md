# 桥接 channel

在某些情况下，你可能会发现自己希望从一系列的chanel 中消费产生的值：
```
<-chan <-chan interface{}
```
这与将channel 切片合并到单个 channel 中稍有不同，如我们在本章前面"The-or-channel"
或 "扇出，扇入"中所看到的。一系列的channel 需要有序地写入，即使是不同的来源。
举一个整个生命周期都在做加法的pipeline stage 的列子。
如果按照我们在本章前面 "约束" 中建立的方式并确保 channel 都被写入它们的 goroutine 拥有,
在每一个新的 goroutine 中的 pipeline stage 重启的时候，都会创建一个新的 channel。
这也就意味着我们很有效地拥有了一个channel系列。
我们将在第5章 "治愈不健康的goroutine" 中详细探讨这种情况。

作为消费者，代码可能不关心其值来自于一系列的channel 的事实。在这种情况下，处理一个充满
channel 的 channel 的可能会很多。如果我们定义一个功能，可以将充满 channel 的 channel
拆解为一个简单的channel（一种称为桥接 channel 的技术），这将使消费者更容易关注手头的问题。
以下是我们如何实现这一目标的一个例子：
```go
bridge := func(
		done <-chan interface{},
		chanStream <-chan <-chan interface{},
	) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				var stream <-chan interface{}
				select {
				case maybeStream, ok := <-chanStream:
					if ok == false {
						return
					}
					stream = maybeStream
				case <-done:
					return
				}
				for val := range orDone(done, stream) {
					select {
					case valStream <- val:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}
```
这是非常简单的代码。现在我们可以使用桥接来实现一个在一个包含多个channel 的 channel上
实现一个单 channel 的门面（facade）。

下面是一个列子，他创建了10个channel，每个channel 都写入一个元素，并将这些channel 
传递给桥接函数:
```go
genVals := func() <-chan <-chan interface {} {
		chanStream := make(chan (<-chan interface{}))
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i ++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}
	for v := range bridge(nil, genVals()) {
		fmt.Printf("%v ", v)
	}
```
这将输出:
```shell
0 1 2 3 4 5 6 7 8 9
```
通过桥接，我们可以在单个range 语句中使用处理 channel 的 channel，
并专注于我们的循环逻辑。将传递 channel 的channel 析构为单一传递数值的channel
来方便别写仅处理数值的逻辑。



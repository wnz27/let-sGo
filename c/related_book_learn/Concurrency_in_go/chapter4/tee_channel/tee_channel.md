# tee-channel
有时候你可能想分割一个来自channel 的值，以便将它们发送到你的代码的两个独立区域中。
设想一下，一个传递用户指令的channel：你可能想要在一个channel上接收一系列用户指令，
将它们发送给相应的执行器，并将它们发送给记录命令以供日后审计的东西。

从类 UNIX 系统中的tee 命令中获得它的名字，tee-channel 就是这样做的。
你可以将它传递给一个读channel，并且它会返回两个单独的channel，以获得相同的值：
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
			for {  // 这里有死循环，所以在有死循环的时候要注意看下有没有从channel 读值并检查 channel 状态
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

	tee := func(
		done <-chan interface{},
		in <-chan interface{},
	) (_, _ <-chan interface{}) {
		out1 := make(chan interface{})
		out2 := make(chan interface{})
		go func() {
			defer close(out1)
			defer close(out2)
			for val := range orDone(done, in) {
				// 我们要使用out1 和 out2 的私有变量版本，所以我们会覆盖这些变量。
				var out1, out2 = out1, out2
				// 我们将使用一条 select 语句，以便不阻塞的写入 out1 和 out2
				// 为确保两者都写入，我们将执行select 语句的两次迭代：每个出站一个channel（没看懂翻译， 看代码理解是每次传入一个）
				for i := 0; i < 2; i ++ {
					select {
					case <-done:
					// 一旦我们写入了channel，我们将其副本设置为nil，以便进一步阻塞写入，而另一个可以继续
					case out1 <- val:
						out1 = nil
					case out2 <- val:
						out2 = nil
					}
				}
			}
		}()
		return out1, out2
	}
}
```

注意写入out1和out2是紧密耦合的，直到out1，和out2 都被写入，迭代才会继续。
通常来说这不是问题，因为处理来自每个channel 的吞吐量都应该是一个确定的某个之外而不是像
tee命令那样，但这并没有任何价值（这句翻译的又没读懂）。下面是一个快速示例
```go
done := make(chan interface{})
	defer close(done)

	out1, out2 := tee(done, take(done, repeat(done, 1, 2), 4))
	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}
```
输出
```shell
out1: 1, out2: 1
out1: 2, out2: 2
out1: 1, out2: 1
out1: 2, out2: 2
```
我们清晰的看到重复利用一个值。
[【Demo】](tee_channel.go)

通过这种模式，对于你的系统来说，继续使用channel 作为 "join点" 将会是易如反掌的事。



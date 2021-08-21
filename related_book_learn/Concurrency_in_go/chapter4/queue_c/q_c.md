# 队列排队
有时，在你的队列尚未准备好的时候就开始接收请求时很有用的。这个过程被称作队列。

这也就意味着只要你的 stage 完成了某些工作，它就会把结果存储在一个稍后其他stage 可以
获取到结果的临时存储位置，而且你的 stage 不需要保存一份指向结果的引用。
在第三章 "channel" 中，我们讨论了带缓存的 channel，那其实就是一种队列，
而且我们当时有足够的理由不去过多讨论使用它。

虽然在系统中引入队列功能非常有用，但它通常是优化程序时希望曹勇的最后一种技术之一。
预先添加队列可以隐藏同步问题，例如死锁和活锁，并且随着程序向正确性收敛，你可能会
发现需要更多或更少的队列。

那么队列有什么好处呢？让我们用一个人们在优化系统性能时经常遇到的一个刦来回答上面的问题：
引入队列来尝试解决性能问题。队列几乎不会加速程序的总运行时间，它只会让程序的行为有所不同。

为了理解它的原因，我们来看看一个简单的pipeline：
```go
done := make(chan interface{})
defer close(done)

zeros := take(done, 3, repeat(done, 0))
short := sleep(done, 1 * time.Second, zeros)
long := sleep(done, 4 * time.Second, short)
pipeline := long
```
这个pipeline 由以下4个 stage 组成：
- 1、一个重复的 stage，会产生层出不穷的0
- 2、当pipeline 中有3个元素的时候会取消上一个 stage，即取消产生"无限个0"的stage。
- 3、暂停1s 的"短" stage。
- 4、一个长的，暂停 4s 的 stage。

对于这个例子来说，我们假设 stage 1 和 stage 2 是即时完成的，让我们关注休眠 stage
如何影响pipeline 的运行时间。

。。。图忽略

这个pipeline 需要大约13s，short stage 大约需要9s来完成。

如果我们修改管道以包含缓冲区会发生什么？让我们来看看对于相同的pipeline，我们在
"长" stage 和 "短" stage 之间引入一个长度为2的缓冲区会发生什么：
```go
done := make(chan interface{})
defer close(done)

zeros := take(done, 3, repeat(done, 0))
short := sleep(done, 1 * time.Second, zeros)
buffer := buffer(done, 2, short) // 给 short stage 进行长度为2 的缓冲
long := sleep(done, 4 * time.Second, short)
pipeline := long
```
运行时间图略。。。看书吧

整个 pipeline 仍然要13s 才能结束运行！但看看 short stage 的运行时间。
它仅在 3s 后就能完成。我们已经将这个stage的运行时间减少了三分之二。
但是如果整个pipeline仍然需要13s来完成执行，增加缓冲又给我们提供了什么帮助呢？

让我们来看下面这个pipeline：
```go
p := processRequest(done, acceptConnection(done, httpHandler))
```
这里pipeline 在取消之前不会退出，并且接收连接的stage 不会停止接收连接，
直到取消channel。在这种情况下，你不希望看到你的程序连接超时，因为你的
processRequest stage 阻止了你的acceptConnection stage。
你希望尽可能地解除你的acceptConnection stage。
否则，你的程序的用户可能会开始发现它们的请求完全被拒绝。

因此，对于引入队列的效用问题的答案并不是一个stage的运行时间已经减少，
而是它处于阻塞状态的时间减少了。这可以让这个stage 继续工作。在这个例子中，
用户可能会在他们的请求中经历滞后，但他们不会被拒绝服务。

通过这种方式，队列的真正用途是将 stage 分离，以便一个 stage 的运行时间不会
影响另一个 stage 的运行时间。以这种方式解耦 stage，然后级联以改变整个系统的
运行时行为，这取决于你的系统，可以是好的也可以是不好的。

然后我们来讨论调整你的排队问题。队列应该放在哪里？缓冲区大小应该是多少？
这些问题的答案取决于你的pipeline 的性质。

首先分析排队可以提高系统整体性能的情况。唯一适用的情况是：
- 如果在一个stage 批处理请求可以节省时间。
- 如果在一个stage 中产生的延迟会在系统中产生一个反馈回环。 
  
### 第一种情况
> 列子是一个 stage 将输入缓存到某些（比如内存）比它本来设计需要发送到介质（例如磁盘）更快的情况。
  
这就是Go 语言 的bufio 包的目的。下面是一个示例，演示了缓冲写入队列与未缓冲写入的简单比较:
```go
package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func BenchmarkUnbufferedWrite(b *testing.B) {
	performWrite(b, tmpFileOrFatal())
}

func BenchmarkBufferedWrite(b *testing.B) {
	bufferedFile := bufio.NewWriter(tmpFileOrFatal())
	performWrite(b, bufio.NewWriter(bufferedFile))
}

func tmpFileOrFatal() *os.File {
	file, err := ioutil.TempFile("", "tmp")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return file
}

func performWrite(b *testing.B, writer io.Writer) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for bt := range take(done, repeat(done, byte(0)), b.N) {
		writer.Write([]byte{bt.(byte)})
	}
}
```
下面运行此基准测试的结果：
```shell
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-4870HQ CPU @ 2.50GHz
BenchmarkUnbufferedWrite-8        200113              5570 ns/op
BenchmarkBufferedWrite-8         1000000              1130 ns/op
PASS
ok      command-line-arguments  2.907s
```
如预期的那样，缓冲写入比未缓冲写入更快。这是因为在bufio.Writer 中，
写入在内部排队到缓冲区， 直到已经积累了足够的块为止，然后块被写出。
这个过程通常称为 分块，原因很明显：
> 分块速度更更快，因为bytes.Buffer 必须增加其分配的内存以容纳它必须存储的字节。
出于各种原因，增长的内存消耗是昂贵的。所以我们需要增长的时间越少，
整个系统的整体效率就越高。因此排队提高了整个系统的性能。

这只是一个简单的内存分块示例，但是你可能会在该领域频繁地进行分块。
通常，任何时候执行操作都需要开销，分块可能会提高系统性能。
这方面的一些例子是打开数据库事务，计算消息校验和以及分配连续空间。

除了分块之外，如果你的算法可以通过支持向后看或排序进行优化，排队起到帮助作用。（这句没看懂！）

### 第二种情况
> 由于一个stage 的延迟导致pipeline 中接收到了更多的输入，
这更难发现，但也更重要，因为它可能导致上游系统的崩溃。

这个想法通常被称为负反馈循环，向下螺旋，甚至是死亡螺旋。这是因为 pipeline 与
上游系统之间存在经常性关系，上游 stage 或系统提交新请求的速度在某种程度上与pipeline 
的有效性有关。

如果pipeline 的效率降低到某个临界阈值以下，在pipeline 上游的系统开始增加它们对
pipeline 的输入，这导致pipeline 损失更多效率，并且死亡螺旋开始。
如果没有某种安全防护，使用pipeline 系统永远不能恢复。

通过在pipeline 入口处引入队列，你可以用创建请求滞后为代价来打破反馈循环。
从调用者进入pipeline 的角度来看，请求似乎正在处理中，但需要很长时间。
只要调用者不超时，你的pipeline 将保持稳定。如果主叫方超时，则需要确保
你在出列时支持某种检查准备情况。
如果你不这样做，你可能会无意中通过处理死亡请求来创建反馈循环，从而降低pipeline 的效率。

#### 你有没有见过死亡螺旋？
> 如果你曾尝试访问一些热门的新系统(例如，新游戏服务器，用于产品发布的网站等)，
> 并且尽管开发人员尽了最大的努力，但该网站一直处于不稳定的状态，恭喜，
> 你可能目睹了一个负面反馈的循环。

> 开发团队开始尝试不同的修复方法，直到有人意识到他们需要一个队列，并且开始匆忙实施。

> 然后客户开会抱怨排队时间。

所以从我们的例子中，我们可以看到一种模式的出现，需要实现排队模式：
- 在你的pipeline 的入口处
- 在这个stage， 批量操作将会带来更高的效率。

你可能会尝试在其他的地方增加队列。例如，在一个重度计算stage 之后。但是请避免
那样的尝试。只有在很少的情况下队列可以减少你的管道的运行时间。而且在队列中胡乱
操作可能会导致灾难性的后果。或许，在最初这并不会显得很明显，我们需要讨论管道的
吞吐量来了解为什么会导致这个情况的原因。

它会帮助我们回答对于如何决定需要用多大的队列的问题。

在"队列" 理论里，有这样的一条法则，通过足够的取样，可以预测pipeline 的需求率。
这被称作利特尔法则。我们首先定义利特尔法则的共识。
它一般被表达为：L = λ * W , 其中：
- L 系统中的平均负载数
- λ 负载的平均到达率
- W 负载在系统中花费的平均时间





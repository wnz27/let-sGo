# 扇出 扇入
假设你已经建立了一条pipeline。
数据流畅地流过你的系统，在你连接在一起的各个 stage 进行转换。
他就像一条美丽的溪流；一个美丽的、缓慢的溪流。

有时候pipeline中某个 stage 可能计算的时间特别耗时。发生这种情况时，你的pipeline中
的上游 stage 可能会被阻塞，同时等待计算耗时的 stage 来完成。
不仅如此，pipeline 本身可能需要很长时间才能全部执行。
> 我们如何解决这个问题呢？

pipeline的一个有趣属性是它们能够让你使用独立的，经常可重新排序的 stage 的组合来操作数据流。
你甚至可以多次重复使用 pipeline 的各个 stage。
在多个 goroutine 上重用我们的 pipeline 的单个stage 以试图并行化来自上游 stage 的pull（push？翻译没懂）
是不是很有趣？也许这将有助于提高 pipeline 的性能。

实际上，事实证明它可以，而这种模式有个名字：扇出，扇入。

> **扇出是一个术语，用于描述启动多个 goroutine 以处理来自pipeline 的输入的过程，
并且扇入是描述将多个结果组合到一个 channel的过程中的术语。**

那么一个什么样的 pipeline 的stage 适合使用这个模式呢？
如果下述两个条件都成立的情况下，你可以考虑在你的stage 中使用扇出模式：
- 他不依赖于之前 stage 计算的值。
- 运行需要很长时间

循序独立性（？隔离，互不相关）很重要，因为你无法保证你的stage 的并发副本以何种顺序运行，
也无法保证其返回的顺序。

我们来看一个例子。在下面的例子中，我构建了一个用来找到素数的非常低效的函数。
我们将使用在本章前面 "pipeline" 中创建的许多 stage。
```go
rand := func() interface{} { return rand.Intn(50000000)}

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := base_func.ToInt(done, base_func.RepeatFn(done, rand))
	fmt.Println("Primes:")
	for prime := range base_func.Take(done, base_func.PrimeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
```
[【demo】我这跟书不太一样](main.go)

输出：
```shell
Primes:
        24941317
        36122539
        6410693
        10128161
        25511527
        2107939
        14004383
        7190363
        45931967
        2393161
Search took: 27.209650425s
```
我们生成一串随机数，最高为50000000， 将数据流传唤为整数流，然后将其传入我们的
PrimeFinder stage。PrimeFinder天真地开始试图将输入流上的数字除以它的每个数字。
如果不成功，它会将该值传递到下一个stage.（似乎没有这个逻辑）
当然这是尝试找到素数的可怕方法，但它符合我们pipeline 将花费 很长的时间的要求。

在我们的循环中，我们通过range 语句循环来遍历出素数，并在找到素数的时候将其输出出来。
多亏了我们的take stage，我们可以在得到10个素数的时候关闭整个pipeline。然后，
我们打印出搜索花了多长时间，并通过defer 语句关闭了完成的 channel，并且将pipeline关闭。

为了避免重复结果出现在我们的结果集中，我们可以在pipeline 中引入另一个 stage 来缓存已在
集合中找到的素数，但为了简单起见，我们将忽略这些。

你可以看到大概花了二十多秒。不是很好。通常我们先看下算法本身。

作为代替，我们将会看看那我们如何扇出一个或者多个stage来让运行缓慢的操作快一些。

这是一个相对简单的例子，而我们只有两个stage：随机数生成和素数筛选。
在更大型的程序中，你的pipeline可能由更多的stage 组成，我们怎么知道该在哪个 stage 扇出？
请记住我们之前的条件：相对顺序独立性和持续时间。
我们的随机数生成器肯定是顺序无关的，因为整数可以是素数也可以不是素数，只有这两种结果，
因为我们所使用的寻找素数的算法，它肯定需要很长时间才能运行。所以他看起来是一个很好的使用
扇出模式的候选。

幸运的是，在pipeline 中分散 stage 的过程非常容易。我们所要做的就是启动多版本的 stage。
我们可以这么做（我这个程序本机原因和书上不太一样，具体可以看书的配到源码，在主目录有贴出来）：
```go
numFinders := runtime.NumCPU()
	finders := make([]<-chan int, numFinders)
	for i := 0; i < numFinders; i ++ {
		finders[i] = base_func.PrimeFinder2(done, randIntStream)
	}
```
[代码见](con_demo/con_1.go)

在这里我们启动了这个 stage 的许多副本，因为我们有多个CPU核心。在我的计算机上，
runtime.NumCPU() 返回8，所以我们将继续在我们的讨论中使用这个数字。
在生产中，我们可能会做一些经验性的测试来确定CPU 的最佳数量，但在这里我们将保持简单，
并且假设只有一个 findPrimes stage 的CPU会被占用。 

我们现在有8个从随机数生成器中取值并试图确定该数字时候是素数的goroutine。
生成随机数不应该花费太多时间，因此 findPrimes stage 的每个goroutine 应该
能够确定它的数字是否为素数，然后立即有另一个随机数可用。（我理解是生产者比消费者快，可以保证这个设想）
但是我们依旧还有一个问题：现在我们有4个goroutine，因此也就输出了4个不同的channel，
但是我们对素数进行迭代的range语句，只能有一个 channel作为输入。
这将是我们使用扇入（fan-in）模式的绝佳实例。

正如我们前面所讨论的，**扇入意味着将多个数据流复用或合并成一个流。**
这样做的算法相对简单：
```go
// FanIn 我们这里采用标准的 done channel 来使我们的 goroutine 可以被关闭，
// 然后用一个可变的interface{} 切片的channel 来进行 扇入
func FanIn(
	done <-chan interface{},
	channels ...<-chan interface{},
) <-chan interface{} {
	// 我们用一个 wait group 来等所有channel 都被处理完
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})
	// 在这里我们创建一个函数，他在传递时将从 channel 中读取，并将读取的值传递到 multiplexedStream channel
	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	// 从所有的 channel 里取值， 往wg 中增加channel 的数量
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	// 等待所有的读操作结束
	go func() {
		// 我们创建一个goroutine 来等待我们多路复用 的 所有channel 被耗尽，
		// 这样我们可以关闭 multiplexedStream channel
		wg.Wait()
		close(multiplexedStream)
	}()
	return multiplexedStream
}
```
[【demo】](base_func/base_func.go)

简而言之，扇入涉及创建用户将读取的多路复用channel，然后为每个传入channel 启动一个
goroutine，以及在传入channel 全部关闭时关闭复用 channel的goroutine。
由于我们要创建一个等待 N个 其他分区完成的 goroutine，创建一个sync.WaitGroup 来
协调是很有意义的。
多路复用功能还通知 WaitGroup 它已完成。

### 额外提醒
> 原生的的扇出扇入算法的实现，仅会在结果到达的顺序不重要的情况。
我们没有做任何事情来保证从 randIntStream 中读取项目的顺序在筛选过程中保留下来。
稍后，我们将看一个保持运行顺序的例子。

让我们把素有这些放一起看看运行时间是否有所减少:
```go
func Con_d_time_dome() {
	done := make(chan interface{})
	defer close(done)
	start := time.Now()

	rand := func() interface{} {
		return rand.Intn(50000000)
	}

	randIntStream := base_func.ToInt(done, base_func.RepeatFn(done, rand))

	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)

	finders := make([]<-chan interface{}, numFinders)
	fmt.Println("Primes:")
	for i := 0; i < numFinders; i++ {
		finders[i] = base_func.PrimeFinder(done, randIntStream)
	}

	for prime := range base_func.Take(done, base_func.FanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
}
```
输出：
```shell
Spinning up 8 prime finders.
Primes:
        6410693
        24941317
        10128161
        36122539
        25511527
        2107939
        14004383
        7190363
        2393161
        45931967
Search took: 6.323565707s
```
所以运行时间从27s降到6s。这清楚地表明了扇出，扇入模式的好处，
它重新定义了pipeline 的用途。我们将执行时间缩短了大约70%多，也不会大幅改变程序的结构。

！！！！太精妙了这个并发模式！！！叹叹叹！！！

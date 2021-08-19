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

> 扇出是一个术语，用于描述启动多个 goroutine 以处理来自pipeline 的输入的过程，
并且扇入是描述将多个结果组合到一个 channel的过程中的术语。

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


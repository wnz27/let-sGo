# 一些便利的生成器

所谓的 "pipeline的生成器" 指的是一个可以把离散值转换为一个channel 上的 数据流的函数。

我们来看一个名为 repeat 的生成器：

```go
package main

func main() {
	repeat := func(
		done <-chan interface{},
		values ...interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}()
		return valueStream
	}
}
```
整个函数会重复你传给它的值，知道你告诉它停止。让我们来看看另一个通用 pipeline stage，
这个 pipeline stage 在复用使用时很有用。参考：
```go
    take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i ++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}
```
这个pipeline stage 只会从其传入的valueStrem 中取出前 num 个 项目，然后退出。
这两个 stage 结合在一起可以非常强大：
```go
    done := make(chan interface{})
	defer close(done)
	
	for num := range take(done, repeat(done, 1), 10) {
		fmt.Printf("%v ", num)
	}
```
以上合起来 [【demo】](repeat_gen/repeat_gen.go)

书上的输出是：
```shell
1 1 1 1 1 1 1 1 1 1 
```
在这个基本的例子中，我们创建了一个重复生成器来生成无限数量的数字1，
但是只取前10个。因为这个"复制生成器"的发送逻辑会因为 "take stage" 没有进行取值而被阻塞，
因此 "重复生成器" 十分高效。尽管我们有能力生成无限数量的流，但我们只生成N + 1 个实例，其中N 是我们
传递到take stage 的数量。

我们可以扩展这一点，让我们创建另一个重复的生成器，但是这次我们创建一个重复调用函数的生成器。
我们称之为repeatFn:
```go
package main

func main() {
	repeatFn := func(
		done <-chan interface{},
		fn func() interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}
}
```
我们用它来生成10个随机数字：
```go
done := make(chan interface{})
	defer close(done)

	rand := func() interface{} {return rand.Int()}

	for num := range take(done, repeatFn(done, rand), 10) {
		fmt.Println(num)
	}
```
输出：
```shell
5577006791947779410
8674665223082153551
6129484611666145821
4037200794235010051
3916589616287113937
6334824724549167320
605394647632969758
1443635317331776148
894385949183117216
2775422040480279449
```
非常酷炫，一个按需生成无限随机整数的 channel！

你可能对于为什么这里的生成器函数以及 stage 都是用 interface{} 作为参数和返回值。
我们可以很容易地将这些函数写成特定的类型，或者可以编写Go语言生成器。

Go语言中的空接口有点不太好，但对于pipeline stage，我认为可以处理 interface{} 的channel，
以便你可以将其视作是一个标准的pipeline 模式的库来使用。正如我们前面所讨论的，
许多pipeline 的效用来自可重用的stage。当 stage 在它所适合的场景的级别下进行操作是最好的。
在 repeat 和 repeatFn 两个生成器中，这两个生成器所关心的是通过列表或某个运算符来生成的一系列数据。

而在 take stage 所关心的仅仅是如何限流我们的pipeline。
这些操作都不关心它们所处理的信息的具体类型，而仅仅需要知道它们参数的数量。

当你需要处理特定的类型时，你可以放置一个为你执行类型断言的stage。有一个专门进行数据断言的额外pipeline stage 
所造成的性能损失可以忽略不计，正如我们稍后会看到的。下面是一个介绍 toString pipeline stage 的小例子:
```go
package main

func main() {
	toString := func(
		done <-chan interface{},
		valueStream <-chan interface{},
	) <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for v := range valueStream {
				select {
				case <-done:
				case stringStream <- v.(string):
				}
			}
		}()
		return stringStream
	}
}
```
以及如何使用它的一个例子：
```go
    done := make(chan interface{})
	defer close(done)
	var message string
	for token := range toString(done, take(done, repeat(done, "I", "am."), 5)) {
		message += token
	}
	fmt.Printf("message: %s...", message)
```
输出：
```shell
message: Iam.Iam.I...
```
[【demo】](gen_assert_type/gen_assert_type.go)

因此让我们向自己证明，将通用部分pipeline 的性能成本忽略不计。
我们将编写两个基准测试函数：一个测试通用stage，一个测试类型特定stage：
```go
package generator

import "testing"

func BenchmarkGeneric(b *testing.B) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for range toString(done, take(done, repeat(done, "a"), b.N)) {

	}
}

func BenchmarkTyped(b *testing.B) {
	repeatT := func(
		done <-chan interface{},
		values ...string,
	) <-chan string {
		valueStream := make(chan string)
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	takeT := func(
		done <-chan interface{},
		valueStream <-chan string,
		num int,
	) <-chan string {
		takeStream := make(chan string)
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i ++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:  // 太他吗容易错了！！！！！
				}
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for range takeT(done, repeatT(done, "a"), b.N) {

	}
}
```
执行以下
```shell
 go test -bench=. -benchtime=1000000x
```
输出为:
```shell
goos: darwin
goarch: amd64
pkg: fzkprac/related_book_learn/Concurrency_in_go/chapter4/generator
cpu: Intel(R) Core(TM) i7-4870HQ CPU @ 2.50GHz
BenchmarkGeneric-8       1000000              1870 ns/op
BenchmarkTyped-8         1000000               967.5 ns/op
PASS
ok      xxxxxxxxxx/xxxxxx 3.029s
```
可以看出来，类型绑定的stage 的速度 是空接口的类型 stage 的两倍，但在整个过程中，也仅仅是稍微快了一点。
一般来说，pipeline 上的限制因素将是你的生成器，或者是计算密集型的一个stage。
【如果生成器不像repeat 和 repeatFn 生成器那样从内存中创建流，则可能会受I/O限制。】
【从磁盘或网络中进行数据读取操作很可能会遮盖住我们这里显示的性能瓶颈。】

如果你的一个 stage 在计算上花费很大，那么这肯定会使这种性能开销消失。（这句不太理解）
如果你认为这个技术依据不符合你的品味的话，你仍然可以通过编写一个Go的生成器来生成你喜欢的数据生成器的stage。

说到一个stage 计算成本很高，我们该如何帮助缓解这个问题呢？它不会限制整个pipeline 的速度吗？
为了缓解这种情况，让我们来讨论扇出扇入（fan-out， fan-in）技术。




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






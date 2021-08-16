# pipeline

当你编写一个程序时，你可能不会坐下来写一个长函数，至少我希望你不要！
你以函数、结构体、方法等形式构造抽象。为什么要这样做呢？
- 其中的部分原因是为了在处理更庞大的任务的时候将细节隐藏起来。
- 另一部分的原因是你可以在不影响其他区域的代码的同时编写某部分的代码。

> 你有没有经历过想要对某个系统进行一处更改，
但你发现你必须更改该系统中的多个区域的代码来达到目的的时候？

这可能是因为你的系统的抽象做的十分糟糕。

pipeline 是你可以用来在系统中形成抽象的另一种工具。特别是，当你的程序需要
流式处理或批处理数据时，它是一个非常强大的工具。

pipeline这个词据称是在1856年首次使用的，可能是指将液体从一个地方输送到另一个地方的
一系列管道。我们在计算机科学中借用了这个术语，因为我们也在从一个地方向另一个地方传输
某些东西：数据！

pipeline只不过是一些列将数据输入，执行操作并将结果数据传回的系统。
【我们称这些操作都是pipeline的一个stage。】

如前所述，一个 stage 只是将数据输入，对其进行转换并将数据送回。这是一个可以被
视为pipeline stage 的函数：
```go
multiply := func(values []int, multiplier int) []int {
		multipliedValues := make([]int, len(values))
		for i, v := range values {
			multipliedValues[i] = v * multiplier
		}
		return multipliedValues
	}
```
这个函数用一个乘法器取一部分整数，随着它的增加循环遍历它们，并返回完成转换之后的
新切片，看起来像一个无聊的函数，我们来创建另一个 stage
```go
add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for i, v := range values {
			addedValues[i] = v + additive
		}
		return addedValues
	}
```
另一个无聊透顶的函数！这个函数式创建一个新的切片，并为每个元素添加一个值。
在这一点上，你可能想知道是什么使这两个函数成为pipeline 的 stage，
而不是仅仅是函数。让我们尝试将它们合并：
```go
ints := []int{1, 2, 3, 4}
	for _, v := range add(multiply(ints, 2), 1) {
		fmt.Println(v)
	}
```
[【批处理的demo】](batch_p/batch_p.go)

输出：
```shell
3
5
7
9
```
看我们是如何在range语句中添加add和multiply函数的。这些函数就像你每天工作的函数一样，
但是因为我们将它们构建为具有pipeline stage 的属性，所以我们可以将它们组合起来形成一个 pipeline。
那很有意思，pipeline stage的属性是什么？
- 一个输入的参数与返回值类型相同的stage。
- 一个stage必须通过编程语言进行 "具体化" 之后才能被当做参数四处传递。
Go语言中的函数就是一种 "具体化"并且很好地贴合我们的需求。
  
> 注：具体化（自己加：我理解是声明）意味着语言向开发人员展示了一个概念，以便它们可以
直接使用它。Go语言中的函数是通用的，因为你可以定义具有函数签名类型的变量。
这也意味着你可以在你的程序中传递函数。

TODO 以下不太懂哈哈哈哈哈
熟悉函数式编程的人可能会点头，并思考像高阶函数和 monad （不懂）这样的术语。

事实上，pipeline stage 与 函数式编程密切相关，可以被认为是 monad 的一个子集。
我不会在这里明确的讨论 monad 或函数式编程，但是它们本身就是一个有趣的主题，并且在尝试
理解pipeline 时，对这两个主题的工作知识虽然不必要，但是有用。

我们的add 和 multiply stage 满足 pipeline stage 的所有属性：
它们都需要一个int类型的切片作为参数并返回一个int类型的切片，并且因为Go语言
具有函数化功能，所以我们可以传递add 和 multiply。

这些属性引起了我们前面提到的pipeline stage 的有趣特性，
即在不改变 stage 本身的情况下，将我们的stage 结合到更高层次变得非常容易。

例如，如果我们现在想要为pipeline添加一个额外的 stage 来乘以2， 我们只需要用
一个新的 multiply stage 包一下之前的pipeline 即可：
```go
ints := []int{1, 2, 3, 4}
for _, v := range multiply(add(multiply(ints, 2), 1), 2) {
	fmt.Println(v)
}
```
输出：
```shell
6
10
14
18
```
注意如何在不编写新函数的前提下修改现有函数或修改我们pipeline的结果。
也许你已经开始看到使用 pipeline 模式的好处了。
当然你也可以编写成下面一样：
```go
ints := []int{1, 2, 3, 4}
for _, v := range ints {
	fmt.Println(2*(v*2 + 1))
}
```
最初这看起来简单得多，但正如我们看到的那样，程序代码在处理【数据流】时不会提供与
pipeline 相同的好处。

请注意每个 stage 是如何获取切片数据并返回切片数据的？
这些stage 正在执行我们称作批处理的操作。
这意味着它们仅大块数据进行一次操作，而不是一次一个离散值。

还有另一种类型的 pipeline stage 执行流处理。 
这意味着这个 stage 一次只接收和处理一个元素。

批处理和流处理有优点和缺点，我们将稍微讨论一下。

**批处理：**
> 现在，请注意，为使原始数据保持不变，每个stage都必须创建一个等长的新片段来存储其计算结果。
这意味我们程序在任何时候的内存占用量都是我们发送到我们pipeline 开始处的切片大小的两倍。

让我们将我们的stage 转换为以流为导向。看起来如下所示
```go
package main

import "fmt"

func main() {
	multiply := func(value int, multiplier int) int {
		return value * multiplier
	}
	add := func(value, additive int) int {
		return value + additive
	}
	ints := []int{1, 2, 3, 4}
	for _, v := range ints {
		fmt.Println(multiply(add(multiply(v, 2), 1), 2))
	}
}
```
[【流处理demo】](stream_p/stream_p.go)
输出：
```shell
6
10
14
18
```
每个stage 都接收并发出一个离散值，我们的程序的内存占用空间将回落到只有pipeline
输入的大小。但是我们不得不将pipeline写入到for 循环之内，
并通过 range 语句为我们的pipeline进行笨拙的提升。
这不仅限制了我们如何向重复利用的pipeline发送消息，而且将在本章后面介绍，
这也限制了我们的扩展能力，还有其他问题。

实际上，我们正在为循环的每次迭代实例化我们的pipeline。
尽管进行函数调用代价很低，但我们为循环的每次迭代进行三次函数调用。

并发性又如何？我前面说过，使用pipeline 的好处之一是能够同时处理各个 stage，
并且我提到了一些关于扇出（fan-out）的内容。那么，它们是从什么地方冒出来的呢？

我本可以通过稍微扩展一下 multiply 和 add 函数来介绍更多相关概念的，
但已经完成了介绍流水线概念的工作。
现在是时候开始学习在Go语言中构建pipeline 的最佳实践了，它始于channel！！



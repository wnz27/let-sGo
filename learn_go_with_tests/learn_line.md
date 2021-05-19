# [来源于 - Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)

重构和工具

本书的重心在于重构的重要性。

好的工具可以帮你放心地进行大型的重构。

你应该足够熟悉你的编辑器，以便使用简单的组合键执行以下操作：

- **提取/内联变量**。能够找出魔法值（magic value）并给他们一个名字可以让你快速简化你的代码。
- **提取方法/功能**。能够获取一段代码并提取函数/方法至关重要。
- **改名**。你应该能够自信地对多个文件内的符号批量重命名。
- **格式化**。Go 有一个名为 go fmt 的专有格式化程序。你的编辑器应该在每次保存文件时都运行它。
- **运行测试**。毫无疑问，你应该能够做到以上任何一点，然后快速重新运行你的测试，以确保你的重构没有破坏任何东西。

另外，为了对你处理代码更有帮助，你应该能够：

- **查看函数签名** - 在 Go 中调用某个函数时，你应该了解并熟悉它。你的 IDE 应根据其文档，参数以及返回的内容描述一个函数。
- **查看函数定义** - 如果仍不清楚函数的功能，你应该能够跳转到源代码并尝试自己弄清楚。
- **查找符号的用法** - 能够查看被调用函数的上下文可以在重构时帮你做出决定。

运用好你的工具将帮助你专注于代码并减少上下文切换。

## 01-hello-world
先编写如下：
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world")
}
```
你打算如何测试这个程序？
将你「领域」内的代码和外界（会引起副作用）分离开会更好。

fmt.Println 会产生副作用（打印到标准输出），我们发送的字符串在自己的领域内。

所以为了更容易测试，我们把这些问题拆分开。
```go
package main

import "fmt"

func Hello() string {
    return "Hello, world"
}

func main() {
    fmt.Println(Hello())
}
```

现在创建一个 hello_test.go 的新文件，来为 Hello 函数编写测试

```go
package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello()
    want := "Hello, world"

    if got != want {
        t.Errorf("got '%q' want '%q'", got, want)
    }
}
```
### 编写测试

编写测试和函数很类似，其中有一些规则
- 程序需要在一个名为 xxx_test.go 的文件中编写
- 测试函数的命名必须以单词 Test 开始
- 测试函数只接受一个参数 t *testing.T
- In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file

> 现在这些信息足以让我们明白，类型为 *testing.T 的变量 t 是你在测试框架中的 hook（钩子），
所以当你想让测试失败时可以执行 t.Fail() 之类的操作。


You can launch the docs locally by running `godoc -http :8000`. 
If you go to localhost:8000/pkg you will see all the packages installed on your system.

The vast majority of the standard library has excellent documentation with examples. 
Navigating to `http://localhost:8000/pkg/testing/` would be worthwhile to see what's available to you.


在上一个示例中，我们在写好代码 之后 编写了测试，以便让你学会如何编写测试和声明函数。从此刻起，我们将 首先编写测试。

#### 我们的下一个需求是指定问候的接受者。

让我们从在测试中捕获这些需求开始。这是基本的测试驱动开发，可以确保我们的测试用例 真正 在测试我们想要的功能。
```go
package main

import "testing"

func TestHelloTo(t *testing.T) {
	got := HelloTo("Alice")
	want := "Hello Alice"
	
	if got != want {
		t.Errorf("got %q want %q", got, want)
    }
}
```
这时需要完成HelloTo
```go
func HelloTo(name string) string {
	return "Hello " + name
}
```
现在再运行测试应该就通过了。通常作为 TDD 周期的一部分，我们该着手 重构 了。

### 关于版本控制的一点说明
> 此时，如果你正在使用版本控制（你应该这样做！）我将按原样 提交 代码。因为我们拥有一个基于测试的可用版本。
不过我不会推送到主分支上，因为我下一步计划重构。现在提交很合适，你总是可以在重构中陷入混乱时回到这个可用版本。

### 常量
这里没有太多可重构的，但我们可以介绍一下另一种语言特性 常量。

通常我们这样定义一个常量
```go
const englishHelloPrefix = "Hello, "
```
现在我们可以重构代码
```go
const englishHelloPrefix = "Hello, "

func Hello(name string) string {
    return englishHelloPrefix + name
}
```
重构之后，重新测试以确保程序无误。

常量应该可以提高应用程序的性能，它避免了每次使用 Hello 时创建 "Hello, " 字符串实例。

显然，对于这个例子来说，性能提升是微不足道的！但是创建常量的价值是可以快速理解值的含义，有时还可以帮助提高性能。

#### 下一个需求是当我们的函数用空字符串调用时，它默认为打印 "Hello, World" 而不是 "Hello, "

首先编写一个新的失败测试
```go

```


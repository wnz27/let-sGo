
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
If you go to `localhost:8000/pkg` you will see all the packages installed on your system.

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

func HelloTo(name string) string {
    return englishHelloPrefix + name
}
```
重构之后，重新测试以确保程序无误。

常量应该可以提高应用程序的性能，它避免了每次使用 Hello 时创建 "Hello, " 字符串实例。

显然，对于这个例子来说，性能提升是微不足道的！但是创建常量的价值是可以快速理解值的含义，有时还可以帮助提高性能。

#### 下一个需求是当我们的函数用空字符串调用时，它默认为打印 "Hello, World" 而不是 "Hello, "

首先编写一个新的失败测试
```go
func TestHello(t *testing.T) {
t.Run("saying hello to people", func(t *testing.T) {
    got := HelloTo("Chris")
    want := "Hello, Chris"
    if got != want {
    t.Errorf("got '%q' want '%q'", got, want)
    }
})

t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
    got := HelloTo("")
    want := "Hello, World"
    
    if got != want {
    t.Errorf("got '%q' want '%q'", got, want)
    }
})

}
```
这里我们将介绍测试库中的另一个工具 -- 子测试。有时，对一个「事情」进行分组测试，然后再对不同场景进行子测试非常有效。

这种方法的好处是，你可以建立在其他测试中也能够使用的共享代码。

当我们检查信息是否符合预期时，会有重复的代码。

重构不 仅仅 是针对程序的代码！
重要的是，你的测试 清楚地说明 了代码需要做什么。

我们可以并且应该重构我们的测试。

```go
func TestHello(t *testing.T) {
    assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got '%q' want '%q'", got, want)
        }
    }

    t.Run("saying hello to people", func(t *testing.T) {
        got := HelloTo("Chris")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })

    t.Run("empty string defaults to 'world'", func(t *testing.T) {
        got := HelloTo("")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })

}
```

t.Helper() 需要告诉测试套件这个方法是辅助函数（helper）。
通过这样做，当测试失败时所报告的行号将在函数调用中而不是在辅助函数内部。
这将帮助其他开发人员更容易地跟踪问题。如果你仍然不理解，请注释掉它，使测试失败并观察测试输出。

来修复代码
```go
const englishHelloPrefix = "Hello, "

func HelloTo(name string) string {
    if name == "" {
        name = "World"
    }
    return englishHelloPrefix + name
}
```
果我运行测试，应该看到它满足了新的要求，并且我们没有意外地破坏其他功能。

回到版本控制
现在我们对代码很满意，我将修改之前的提交，所以我们只提交认为好的版本及其测试。

### 规律
让我们再次回顾一下这个周期

- 编写一个测试
- 让编译通过
- 运行测试，查看失败原因并检查错误消息是很有意义的
- 编写足够的代码以使测试通过
- 重构
从表面上看可能很乏味，但坚持这种反馈循环非常重要。

它不仅确保你有 相关的测试，还可以确保你通过重构测试的安全性来 设计优秀的软件。

查看测试失败是一个重要的检查手段，因为它还可以让你看到错误信息。
作为一名开发人员，如果测试失败时不能清楚地说明问题所在，那么使用这个代码库可能会非常困难。

通过确保你的测试的 快速，并设置你的工具，可以使运行测试足够简单，你在编写代码时就可以进入流畅的状态。

如果不写测试，你提交的时候通过运行软件来手动检查你的代码，这会打破你的流畅状态，
而且你任何时候都无法将自己从这种状态中拯救出来，尤其是从长远来看。

### 更多需求

现在需要支持第二个参数，指定问候的语言。如果一种不能识别的语言被传进来，就默认为英语。

为使用西班牙语的用户编写测试，将其添加到现有的测试用例中。
```go
t.Run("in Spanish", func(t *testing.T) {
    got := Hello("Elodie", "Spanish")
    want := "Hola, Elodie"
    assertCorrectMessage(t, got, want)
})
```
先编写测试。当你尝试运行测试时，编译器 应该 会出错，因为你用两个参数而不是一个来调用 HelloTo
```go
./hello_test.go:41:17: too many arguments in call to HelloTo
        have (string, string)
        want (string)
```

通过向 Hello 添加另一个字符串参数来解决编译问题
```go
func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }
    return englishHelloPrefix + name
```
当你尝试再次运行测试时，它会报错在其他测试和 hello.go 中没有传递足够的参数给 Hello 函数
```go
./hello_test.go:35:17: not enough arguments in call to HelloTo
        have (string)
        want (string, string)
```
通过传递空字符串来解决它们。现在，除了我们的新场景外，你的所有测试都应该编译并通过
```go
--- FAIL: TestHelloTo (0.00s)
    --- FAIL: TestHelloTo/in_Spanish (0.00s)
        hello_test.go:43: got "Hello Elodie" want "Hola, Elodie"
```
这里我们可以使用 if 检查语言是否是「西班牙语」，如果是就修改信息
```go
func HelloTo(name string, language string) string {
	if name == "" {
		name = "world"
	}
	if language == "Spanish" {
		return "Hola " + name
	}
	return englishHelloPrefix + name
}
```
现在是 重构 的时候了。 你应该在代码中看出了一些问题，其中有一些重复的「魔术」字符串。
自己尝试重构它，每次更改都要重新运行测试，以确保重构不会破坏任何内容。
```go
const spanish = "Spanish"
const spanishHelloPrefix = "Hola "
const englishHelloPrefix = "Hello "

func HelloTo(name string, language string) string {
	if name == "" {
		name = "world"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}
```
#### 法语

编写一个测试，断言如果你传递 "French" 你会得到 "Bonjour, "
看到它失败，检查错误信息是否容易理解
在代码中进行最小的合理更改

可能如下
```go
func HelloTo(name string, language string) string {
	if name == "" {
		name = "world"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	if language == french {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}
```
#### switch
当你有很多 if 语句检查一个特定的值时，通常使用 switch 语句来代替。
如果我们希望稍后添加更多的语言支持，我们可以使用 switch 来重构代码，使代码更易于阅读和扩展。
```go
func HelloTo(name string, language string) string {
	if name == "" {
		name = "world"
	}
	prefix := englishHelloPrefix
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}
	return prefix + name
}
```
编写一个测试，添加用你选择的语言写的问候，你应该可以看到扩展这个 神奇 的函数是多么简单。

### 最后一次重构？
你可能会抱怨说也许我们的函数正在变得很臃肿。对此最简单的重构是将一些功能提取到另一个函数中。
```go
func HelloTo(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string)  (prefix string){
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
```
一些新的概念：

- 在我们的函数签名中，我们使用了 命名返回值（prefix string）。
- 这将在你的函数中创建一个名为 prefix 的变量。
  - 它将被分配「零」值。这取决于类型，例如 int 是 0，对于字符串它是 ""。
    - 你只需调用 return 而不是 return prefix 即可返回所设置的值。 
  - 这将显示在 Go Doc 中，所以它使你的代码更加清晰。
- 如果没有其他 case 语句匹配，将会执行 default 分支。
- 函数名称以小写字母开头。在 Go 中，公共函数以大写字母开始，私有函数以小写字母开头。
  我们不希望我们算法的内部结构暴露给外部，所以我们将这个功能私有化。


### Go 的一些语法

- 编写测试
- 用参数和返回类型声明函数
    if，const，switch
- 声明变量和常量

### TDD 过程以及步骤的重要性

- 编写一个失败的测试，并查看失败信息，我们知道现在有一个为需求编写的 相关 的测试，并且看到它产生了 易于理解的失败描述 
- 编写最少量的代码使其通过，以获得可以运行的程序
- 然后 重构，基于我们测试的安全性，以确保我们拥有易于使用的精心编写的代码

在我们的例子中，我们通过小巧易懂的步骤从 Hello() 到 Hello("name")，到 Hello("name", "french")。

与「现实世界」的软件相比，这当然是微不足道的，但原则依然通用。
TDD 是一门需要通过开发去实践的技能，通过将问题分解成更小的可测试的组件，你编写软件将会更加轻松。

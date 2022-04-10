# 迭代

在 Go 中 for 用来循环和迭代，Go 语言没有 while，do，until 这几个关键字，你只能使用 for。这也算是件好事！

让我们来为一个重复字符 5 次的函数编写测试。

目前这里没什么新知识，所以你可以自己尝试去写

## 先写测试
```go
func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expeted := "aaaaa"

	if repeated != expeted {
		t.Errorf("expected %q but got %q", expeted, repeated)
	}
}

```

运行测试`./repeat_test.go:6:14: undefined: Repeat`

## 先使用最少的代码来让失败的测试先跑起来

请遵守原则！你现在不需要学习任何新知识就可以让测试恰当地失败。

现在只需让代码可编译，这样你就可以检查测试用例能否通过。
```go
package iteration
func Repeat(character string) string {
    return ""
}
```

现在你已经掌握了足够的 Go 知识来给一些基本的问题编写测试，
这意味着你可以放心的处理生产环境的代码，并知道它的行为会如你所愿。

`repeat_test.go:10: expected 'aaaaa' but got '`

## 把代码补充完整，使得它能够通过测试

就像大多数类 C 的语言一样，for 语法很不起眼。
```go
func Repeat(character string) string {
    var repeated string
    for i := 0; i < 5; i++ {
        repeated = repeated + character
    }
    return repeated
}
```
与其它语言如 C，Java 或 JavaScript 不同，在 Go 中 for 语句前导条件部分并没有圆括号，而且大括号 { } 是必须的。
你可能会好奇下面这行
```go
var repeated string
```
我们目前都是使用 := 来声明和初始化变量。然后 := 只是两个步骤的简写。
这里我们使用显式的版本来声明一个 string 类型的变量。我们还可以使用 var 来声明函数，稍后我们将看到这一点。
运行测试应该是通过的。

关于 for 循环其它变体请参考[这里](https://gobyexample.com/for)。

## 重构

现在是时候重构并引入另一个构造（construct）+= 赋值运算符。
```go
const repeatCount = 5
func Repeat(character string) string {
    var repeated string
    for i := 0; i < repeatCount; i++ {
        repeated += character
    }
    return repeated
}
```
+= 是自增赋值运算符（Add AND assignment operator），它把运算符右边的值加到左边并重新赋值给左边。
它在其它类型也可以使用，比如整数类型。

## 基准测试

在 Go 中编写 [基准测试（benchmarks）](https://golang.org/pkg/testing/#hdr-Benchmarks) 是该语言的另一个一级特性，它与编写测试非常相似。

```go
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}
```
你会看到上面的代码和写测试差不多。

testing.B 可使你访问隐性命名（cryptically named）b.N。

基准测试运行时，代码会运行 b.N 次，并测量需要多长时间。

代码运行的次数不会对你产生影响，测试框架会选择一个**它所认为的最佳值**，以便让你获得更合理的结果。

用 `go test -bench=.` 来运行基准测试。 (如果在 Windows Powershell 环境下使用 `go test -bench="."`)
```go
goos: darwin
goarch: amd64
pkg: fzkprac/learn_go_with_tests/03-iterations
cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
BenchmarkRepeat-4        7983950               132.0 ns/op
PASS
ok      fzkprac/learn_go_with_tests/03-iterations       1.953s
```
以上结果说明运行一次这个函数需要 132 纳秒（在我的电脑上）。这挺不错的，为了测试它运行了 10000000 次。

注意：基准测试默认是顺序运行的。

## 练习
修改测试代码，以便调用者可以指定字符重复的次数，然后修复代码
写一个 ExampleRepeat 来完善你的函数文档
看一下 [strings](https://golang.org/pkg/strings/) 包。找到你认为可能有用的函数，并对它们编写一些测试。
投入时间学习标准库会慢慢得到回报。

## 总结

更多的 TDD 练习
学习了 for 循环
学习了如何编写基准测试



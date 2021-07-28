## 简介
反射是一种机制，在编译时不知道具体类型的情况下，可以透视结构的组成、更新值。使用反射，可以让我们编写出能统一处理所有类型的代码。甚至是编写这部分代码时还不存在的类型。一个具体的例子就是fmt.Println()方法，可以打印出我们自定义的结构类型。

虽然，一般来说都不建议在代码中使用反射。反射影响性能、不易阅读、将编译时就能检查出来的类型问题推迟到运行时以 panic 形式表现出来，这些都是反射的缺点。但是，我认为反射是一定要掌握的，原因如下：

- 很多标准库和第三方库都用到了反射，虽然暴露的接口做了封装，不需要了解反射。但是如果要深入研究这些库，了解实现，阅读源码， 反射是绕不过去的。例如encoding/json，encoding/xml等；
- 如果有一个需求，编写一个可以处理所有类型的函数或方法，我们就必须会用到反射。因为 Go 的类型数量是无限的，而且可以自定义类型，所以使用类型断言是无法达成目标的。
Go 语言标准库reflect提供了反射功能。

## 接口
反射是建立在 Go 的类型系统之上的，并且与接口密切相关。
首先简单介绍一下接口。
Go 语言中的接口约定了一组方法集合，
任何定义了这组方法的类型（也称为实现了接口）的变量都可以赋值给该接口的变量。
[demo](./interface_demo1/d1.go)
```go
package main

import "fmt"

type Animal interface {
	Speak()
}

type Cat struct {
}

func (c Cat) Speak() {
	fmt.Println("Meow")
}

type Dog struct {
}

func (d Dog) Speak() {
	fmt.Println("Bark")
}

func main() {
	var a Animal

	a = Cat{}
	a.Speak()
	fmt.Printf("%v\n", &a)

	a = Dog{}
	a.Speak()
	fmt.Printf("%v\n", &a)
}
```
上面代码中，我们定义了一个Animal接口，它约定了一个方法Speak()。
而后定义了两个结构类型Cat和Dog，都定义了这个方法。
这样，我们就可以将Cat和Dog对象赋值给Animal类型的变量了。



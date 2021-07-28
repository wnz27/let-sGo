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

接口变量包含两部分：类型和值，即(type, value)。
类型就是赋值给接口变量的值的类型，值就是赋值给接口变量的值。
如果知道接口中存储的变量类型，我们也可以使用类型断言通过接口变量获取具体类型的值：
```go
type Animal interface {
  Speak()
}

type Cat struct {
  Name string
}

func (c Cat) Speak() {
  fmt.Println("Meow")
}

func main() {
  var a Animal

  a = Cat{Name: "kitty"}
  a.Speak()

  c := a.(Cat)
  fmt.Println(c.Name)
}
```
上面代码中，我们知道接口a中保存的是Cat对象，直接使用类型断言a.(Cat)获取Cat对象。
但是，如果类型断言的类型与实际存储的类型不符，会直接 panic。
所以实际开发中，通常使用另一种类型断言形式`c, ok := a.(Cat)`。
如果类型不符，这种形式不会 panic，而是通过将第二个返回值置为 false 来表明这种情况。

有时候，一个类型定义了很多方法，而不只是接口约定的方法。
通过接口，我们只能调用接口中约定的方法。当然我们也可以将其类型断言为另一个接口，
然后调用这个接口约定的方法，前提是原对象实现了这个接口：
```go
var r io.Reader
r = new(bytes.Buffer)
w = r.(io.Writer)
```
io.Reader和io.Writer是标准库中使用最为频繁的两个接口：
```go
// src/io/io.go
type Reader interface {
  Read(p []byte) (n int, err error)
}
type Writer interface {
  Write(p []byte) (n int, err error)
}
```

bytes.Buffer同时实现了这两个接口，所以byte.Buffer对象可以赋值给io.Reader变量r，
然后r可以断言为io.Writer，因为接口io.Reader中存储的值也实现了io.Writer接口。

如果一个接口A包含另一个接口B的所有方法，那么接口A的变量可以直接赋值给B的变量，
因为A中存储的值一定实现了A约定的所有方法，那么肯定也实现了B。
此时，无须类型断言。
例如标准库io中还定义了一个io.ReadCloser接口，此接口变量可以直接赋值给io.Reader：
```go
// src/io/io.go
type ReadCloser interface {
  Reader
  Closer
}
```
空接口interface{}是比较特殊的一个接口，它没有约定任何方法。
所有类型值都可以赋值给空接口类型的变量，因为它没有任何方法限制。
有一点特别重要，接口变量之间类型断言也好，直接赋值也好，其内部存储的(type, value)类型-值对是没有变化的。
只是通过不同的接口能调用的方法有所不同而已。
也是由于这个原因，接口变量中存储的值一定不是接口类型。

## 反射基础





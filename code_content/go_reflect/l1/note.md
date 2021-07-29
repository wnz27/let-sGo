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
Go 语言中的反射功能由reflect包提供。

reflect包定义了一个接口reflect.Type和一个结构体reflect.Value，
它们定义了大量的方法用于获取类型信息，设置值等。

在reflect包内部，只有类型描述符实现了reflect.Type接口。
由于类型描述符是未导出类型，我们只能通过reflect.TypeOf()方法获取reflect.Type类型的值：
```go
package main

import (
  "fmt"
  "reflect"
)

type Cat struct {
  Name string
}

func main() {
  var f float64 = 3.5
  t1 := reflect.TypeOf(f)
  fmt.Println(t1.String())

  c := Cat{Name: "kitty"}
  t2 := reflect.TypeOf(c)
  fmt.Println(t2.String())
}
```
输出：
```go
float64
main.Cat
```
Go 语言是静态类型的，每个变量在编译期有且只能有一个确定的、已知的类型，即变量的静态类型。
静态类型在变量声明的时候就已经确定了，无法修改。
一个接口变量，它的静态类型就是该接口类型。

虽然在运行时可以将不同类型的值赋值给它，改变的也只是它内部的动态类型和动态值。
它的静态类型始终没有改变。

reflect.TypeOf()方法就是用来取出接口中的动态类型部分，以reflect.Type返回。

上面代码好像并没有接口类型啊？
我们看下reflect.TypeOf()的定义：
```go
// src/reflect/type.go
func TypeOf(i interface{}) Type {
  eface := *(*emptyInterface)(unsafe.Pointer(&i))
  return toType(eface.typ)
}
```
它接受一个interface{}类型的参数，所以上面的float64和Cat变量会先转为interface{}再传给方法，
reflect.TypeOf()方法获取的就是这个interface{}中的类型部分。

相应地，reflect.ValueOf()方法自然就是获取接口中的值部分，返回值为reflect.Value类型。
在上例基础上添加下面代码：
```go
v1 := reflect.ValueOf(f)
fmt.Println(v1)
fmt.Println(v1.String())

v2 := reflect.ValueOf(c)
fmt.Println(v2)
fmt.Println(v2.String())
```
输出：
```go
3.5
<float64 Value>
{kitty}
<main.Cat Value>
```
由于fmt.Println()会对reflect.Value类型做特殊处理，打印其内部的值，
所以上面显示调用了reflect.Value.String()方法获取更多信息。

获取类型如此常见，fmt提供了格式化符号%T输出参数类型：
```go
fmt.Printf("%T\n", 3) // int
```
Go 语言中类型是无限的，而且可以通过type定义新的类型。
但是类型的种类是有限的，reflect包中定义了所有种类的枚举：
```go
// src/reflect/type.go
type Kind uint
const (
  Invalid Kind = iota
  Bool
  Int
  Int8
  Int16
  Int32
  Int64
  Uint
  Uint8
  Uint16
  Uint32
  Uint64
  Uintptr
  Float32
  Float64
  Complex64
  Complex128
  Array
  Chan
  Func
  Interface
  Map
  Ptr
  Slice
  String
  Struct
  UnsafePointer
)
```
一共 26 种，我们可以分类如下：
- 基础类型Bool、String以及各种数值类型
    - 有符号整数Int/Int8/Int16/Int32/Int64
    - 无符号整数Uint/Uint8/Uint16/Uint32/Uint64/Uintptr
    - 浮点数Float32/Float64
    - 复数Complex64/Complex128
- 复合（聚合）类型Array和Struct
- 引用类型Chan、Func、Ptr、Slice和Map（值类型和引用类型区分不明显，这里不引战，大家理解意思就行）
- 接口类型Interface
- 非法类型Invalid，表示它还没有任何值（reflect.Value的零值就是Invalid类型）

Go 中所有的类型（包括自定义的类型），都是上面这些类型或它们的组合。
例如：
```go
type MyInt int

func main() {
  var i int
  var j MyInt

  i = int(j) // 必须强转

  ti := reflect.TypeOf(i)
  fmt.Println("type of i:", ti.String())

  tj := reflect.TypeOf(j)
  fmt.Println("type of j:", tj.String())

  fmt.Println("kind of i:", ti.Kind())
  fmt.Println("kind of j:", tj.Kind())
}
```
上面两个变量的静态类型分别为int和MyInt，是不同的。
虽然MyInt的底层类型（underlying type）也是int。
它们之间的赋值必须要强制类型转换。

但是它们的种类是一样的，都是int。
代码输出如下：
```go
type of i: int
type of j: main.MyInt
kind of i: int
kind of j: int
```

## 反射用法
结合具体用法来说

### 透视数据组成
透视结构体组成，需要以下方法：
- reflect.ValueOf()：获取反射值对象；
- reflect.Value.NumField()：从结构体的反射值对象中获取它的字段个数；
- reflect.Value.Field(i)：从结构体的反射值对象中获取第i个字段的反射值对象；
- reflect.Kind()：从反射值对象中获取种类；
- reflect.Int()/reflect.Uint()/reflect.String()/reflect.Bool()：这些方法从反射值对象做取出具体类型。

示例
```go
type User struct {
  Name    string
  Age     int
  Married bool
}

func inspectStruct(u interface{}) {
  v := reflect.ValueOf(u)
  for i := 0; i < v.NumField(); i++ {
    field := v.Field(i)
    switch field.Kind() {
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
      fmt.Printf("field:%d type:%s value:%d\n", i, field.Type().Name(), field.Int())

    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
      fmt.Printf("field:%d type:%s value:%d\n", i, field.Type().Name(), field.Uint())

    case reflect.Bool:
      fmt.Printf("field:%d type:%s value:%t\n", i, field.Type().Name(), field.Bool())

    case reflect.String:
      fmt.Printf("field:%d type:%s value:%q\n", i, field.Type().Name(), field.String())

    default:
      fmt.Printf("field:%d unhandled kind:%s\n", i, field.Kind())
    }
  }
}

func main() {
  u := User{
    Name:    "dj",
    Age:     18,
    Married: true,
  }

  inspectStruct(u)
}
```
结合使用reflect.Value的NumField()和Field()方法可以遍历结构体的每个字段。
然后针对每个字段的Kind做相应的处理。

有些方法只有在原对象是某种特定类型时，才能调用。
例如NumField()和Field()方法只有原对象是结构体时才能调用，否则会panic。

识别出具体类型后，可以调用反射值对象的对应类型方法获取具体类型的值，例如上面的field.Int()/field.Uint()/field.Bool()/field.String()。
但是为了减轻处理的负担，Int()/Uint()方法对种类做了合并处理，它们只返回相应的最大范围的类型，Int()返回Int64类型，Uint()返回Uint64类型。
而Int()/Uint()内部会对相应的有符号或无符号种类做处理，转为Int64/Uint64返回。下面是reflect.Value.Int()方法的实现：
```go
// src/reflect/value.go
func (v Value) Int() int64 {
  k := v.kind()
  p := v.ptr
  switch k {
  case Int:
    return int64(*(*int)(p))
  case Int8:
    return int64(*(*int8)(p))
  case Int16:
    return int64(*(*int16)(p))
  case Int32:
    return int64(*(*int32)(p))
  case Int64:
    return *(*int64)(p)
  }
  panic(&ValueError{"reflect.Value.Int", v.kind()})
}
```
上面代码，我们只处理了少部分种类。在实际开发中，
完善的处理需要破费一番功夫，特别是字段是其他复杂类型，甚至包含循环引用的时候。

#### 另外，我们也可以透视标准库中的结构体，并且可以透视其中的未导出字段。
使用上面定义的inspectStruct()方法：
```go
inspectStruct(bytes.Buffer{})
```
bytes.Buffer的结构如下：
```go
type Buffer struct {
  buf      []byte
  off      int   
  lastRead readOp
}
```
都是未导出的字段，程序输出：
```go
field:0 unhandled kind:slice
field:1 type:int value:0
field:2 type:readOp value:0
```

#### 透视map组成，需要以下方法：
- reflect.Value.MapKeys()：将每个键的reflect.Value对象组成一个切片返回；
- reflect.Value.MapIndex(k)：传入键的reflect.Value对象，返回值的reflect.Value；
- 然后可以对键和值的reflect.Value进行和上面一样的处理。

示例：
```go
func inspectMap(m interface{}) {
  v := reflect.ValueOf(m)
  for _, k := range v.MapKeys() {
    field := v.MapIndex(k)

    fmt.Printf("%v => %v\n", k.Interface(), field.Interface())
  }
}

func main() {
  inspectMap(map[uint32]uint32{
    1: 2,
    3: 4,
  })
}
```
我这里偷懒了，没有针对每个Kind去做处理，直接调用键-值reflect.Value的Interface()方法。
该方法以空接口的形式返回内部包含的值。程序输出：
```go
1 => 2
3 => 4
```
同样地，MapKeys()和MapIndex(k)方法只能在原对象是map类型时才能调用，否则会panic。

#### 透视切片或数组组成，需要以下方法：
- reflect.Value.Len()：返回数组或切片的长度；
- reflect.Value.Index(i)：返回第i个元素的reflect.Value值；
- 然后对这个reflect.Value判断Kind()进行处理。
示例：
```go
func inspectSliceArray(sa interface{}) {
  v := reflect.ValueOf(sa)

  fmt.Printf("%c", '[')
  for i := 0; i < v.Len(); i++ {
    elem := v.Index(i)
    fmt.Printf("%v ", elem.Interface())
  }
  fmt.Printf("%c\n", ']')
}

func main() {
  inspectSliceArray([]int{1, 2, 3})
  inspectSliceArray([3]int{4, 5, 6})
}
```
同样地Len()和Index(i)方法只能在原对象是切片，
数组或字符串时才能调用，其他类型会panic。

#### 透视函数类型，需要以下方法：

reflect.Type.NumIn()：获取函数参数个数；
reflect.Type.In(i)：获取第i个参数的reflect.Type；
reflect.Type.NumOut()：获取函数返回值个数；
reflect.Type.Out(i)：获取第i个返回值的reflect.Type。
示例：
```go
func Add(a, b int) int {
  return a + b
}

func Greeting(name string) string {
  return "hello " + name
}

func inspectFunc(name string, f interface{}) {
  t := reflect.TypeOf(f)
  fmt.Println(name, "input:")
  for i := 0; i < t.NumIn(); i++ {
    t := t.In(i)
    fmt.Print(t.Name())
    fmt.Print(" ")
  }
  fmt.Println()

  fmt.Println("output:")
  for i := 0; i < t.NumOut(); i++ {
    t := t.Out(i)
    fmt.Print(t.Name())
    fmt.Print(" ")
  }
  fmt.Println("\n===========")
}

func main() {
  inspectFunc("Add", Add)
  inspectFunc("Greeting", Greeting)
}
```
同样地，只有在原对象是函数类型的时候才能调用NumIn()/In()/NumOut()/Out()这些方法，其他类型会panic。

程序输出：
```go
Add input:
int int
output:
int
===========
Greeting input:
string
output:
string
===========
```
#### 透视结构体中定义的方法，需要以下方法：
- reflect.Type.NumMethod()：返回结构体定义的方法个数；
- reflect.Type.Method(i)：返回第i个方法的reflect.Method对象；

示例：
```go
func inspectMethod(o interface{}) {
  t := reflect.TypeOf(o)

  for i := 0; i < t.NumMethod(); i++ {
    m := t.Method(i)

    fmt.Println(m)
  }
}

type User struct {
  Name    string
  Age     int
}

func (u *User) SetName(n string) {
  u.Name = n
}

func (u *User) SetAge(a int) {
  u.Age = a
}

func main() {
  u := User{
    Name:    "dj",
    Age:     18,
  }
  inspectMethod(&u)
}
```
reflect.Method定义如下：
```go
// src/reflect/type.go
type Method struct {
  Name    string // 方法名
  PkgPath string

  Type  Type  // 方法类型（即函数类型）
  Func  Value // 方法值（以接收器作为第一个参数）
  Index int   // 是结构体中的第几个方法
}
```
事实上，reflect.Value也定义了NumMethod()/Method(i)这些方法。
区别在于：reflect.Type.Method(i)返回的是一个reflect.Method对象，
可以获取方法名、类型、是结构体中的第几个方法等信息。
如果要通过这个reflect.Method调用方法，必须使用Func字段，
而且要传入接收器的reflect.Value作为第一个参数：
```go
m.Func.Call(v, ...args)
```
但是reflect.Value.Method(i)返回一个reflect.Value对象，
它总是以调用Method(i)方法的reflect.Value作为接收器对象，不需要额外传入。
而且直接使用Call()发起方法调用：
```go
m.Call(...args)
```
reflect.Type和reflect.Value有不少同名方法，使用时需要注意甄别。
## 调用函数或方法
调用函数，需要以下方法：
reflect.Value.Call()：使用reflect.ValueOf()生成每个参数的反射值对象，然后组成切片传给Call()方法。
Call()方法执行函数调用，返回[]reflect.Value。其中每个元素都是原返回值的反射值对象。

示例
```go
package main

import (
	"fmt"
	"reflect"
)

func Add(a, b int) int {
	return a + b
}

func Greeting(name string) string {
	return "hello " + name
}

func invoke(f interface{}, args ...interface{}) {
	v := reflect.ValueOf(f)

	argsV := make([]reflect.Value, 0, len(args))
	for _, arg := range args {
		argsV = append(argsV, reflect.ValueOf(arg))
	}

	rets := v.Call(argsV)

	fmt.Println("ret:")
	for _, ret := range rets {
		fmt.Println(ret.Interface())
	}
}

func main() {
	invoke(Add, 1, 2)
	invoke(Greeting, "dj")
}

```
我们封装一个invoke()方法，以interface{}空接口接收函数对象，
以interface{}可变参数接收函数调用的参数。
函数内部首先调用reflect.ValueOf()方法获得函数对象的反射值对象。

然后依次对每个参数调用reflect.ValueOf()，生成参数的反射值对象切片。
最后调用函数反射值对象的Call()方法，输出返回值。

程序运行结果：
```go
ret:
3
ret:
hello dj
```
方法的调用也是类似的：
```go
type M struct {
  a, b int
  op   rune
}

func (m M) Op() int {
  switch m.op {
  case '+':
    return m.a + m.b

  case '-':
    return m.a - m.b

  case '*':
    return m.a * m.b

  case '/':
    return m.a / m.b

  default:
    panic("invalid op")
  }
}

func main() {
  m1 := M{1, 2, '+'}
  m2 := M{3, 4, '-'}
  m3 := M{5, 6, '*'}
  m4 := M{8, 2, '/'}
  invoke(m1.Op)
  invoke(m2.Op)
  invoke(m3.Op)
  invoke(m4.Op)
}
```
运行结果
```go
ret:
3
ret:
-1
ret:
30
ret:
4
```
以上是在编译期明确知道方法名的情况下发起调用。
如果只给一个结构体对象，通过参数指定具体调用哪个方法该怎么做呢？这需要以下方法：
- reflect.Value.MethodByName(name)：获取结构体中定义的名为name的方法的reflect.Value对象，
  这个方法默认有接收器参数，即调用MethodByName()方法的reflect.Value。

示例：
```go

```
我们可以在结构体的反射值对象上使用NumMethod()和Method()遍历它定义的所有方法。


## 实战案例



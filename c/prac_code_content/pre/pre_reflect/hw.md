反射是Go里面很重要的一个特性，反射是程序在运行时检查其变量和值并找到它们类型的能力。你可能不明白这意味着什么，但这没关系。通过这个作业你将对反射有一个清晰的了解。

通过这个作业我希望大家了解这些概念

- reflect.Type和 reflect.Value
- reflect.Kind
- NumField() 和 Field() 方法
- Int() 和 String()

通过下面的例子来动手实战

提供下面两个struct结构体

```go
type order struct {
ordId      int
customerId int
}

type employee struct {
name    string
id      int
address string
salary  int
country string
}
```

下面是主函数的实现，就是通过传入的struct来生成SQL语句， 加入我们要根据不同的结构体插入到不同的数据库表里面去。

```go
func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)
	i := 90
	createQuery(i)
}
```

实现createQuery

```go
func createQuery(q interface{}) {
	
}
```

最终的输出是这三个

```shell
insert into order values(456, 56)  
insert into employee values("Naveen", 565, "Coimbatore", 90000, "India")  
unsupported type
```

[【hw --- demo】](pre_reflect.go)

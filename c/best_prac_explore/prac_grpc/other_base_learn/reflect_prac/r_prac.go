/**
 * @project prac_grpc
 * @Author 27
 * @Description //TODO
 * @Date 2021/6/6 12:03 6月
 **/
package main

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
)

type HelloService interface {
SayHello(name string) (string, error)
GetOrder(name string) (string, error)
}

// golang 常用一种写法，用于确保肯定实现了某个接口
var _ HelloService = hello{}

func Hello() *hello {
	return &hello{endpoint: "http://localhost:8080/"}
}

type hello struct {
	endpoint string
	FuncField func()
}

func (h hello) GetOrder(name string) (string, error) {
	return "mock", nil
}

func (h hello) SayHello(name string) (string, error) {
	client := http.Client{}
	resp, err := client.Get(h.endpoint + name)
	if err != nil {
		//fmt.Printf("%+v", err.Error())
		return "", err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		//fmt.Printf("%+v", err)
		return "", err
	}

	fmt.Printf("%v", data)

	return string(data), nil
}

func PrintFuncName(val interface{}) {
	// 反射 reflection
	t := reflect.TypeOf(val)
	num := t.NumMethod()
	for i := 0; i < num; i ++ {
		m := t.Method(i)
		fmt.Println(m.Name)

	}
}


func PrintFieldName(val interface{}) {
	// 反射 reflection
	t := reflect.TypeOf(val)
	num := t.NumField()
	for i := 0; i < num; i ++ {
		f := t.Field(i)
		fmt.Println(f.Name)
	}
}


func PrintTypeValue(val interface{}){
	t := reflect.TypeOf(val)
	v := reflect.ValueOf(val)

	numField := t.NumField()
	for i := 0; i < numField; i++ {
		field := t.Field(i)
		fmt.Println(field.Name)
		fieldValue := v.Field(i)
		fmt.Println(fieldValue.CanSet())
		if fieldValue.CanSet() {
			fmt.Println("ddd")
		}
	}
}
/*
// CanSet reports whether the value of v can be changed.
// A Value can be changed only if it is addressable and was not
// obtained by the use of unexported struct fields.
// If CanSet returns false, calling Set or any type-specific
// setter (e.g., SetBool, SetInt) will panic.
 */
// !  addressable 可寻址的，意味着指针！！！

/*
Golang 语法 —— 方法接收器用哪个?
• 核心原则:遇事不决用指针
• 次级原则:不可变对象用结构体

注:早期入门，不要在这个点纠缠。拿捏不定的时 候，写个测试就知道该用啥
 */

// 为指针写的方法
func PrintFieldNameFroPtr(val interface{}) {
	v := reflect.ValueOf(val)  // 这是指针的反射
	ele := v.Elem()  // 拿到了指针指向的结构体
	t := ele.Type()  // 拿到了指针指向的结构体的类型信息

	numField := t.NumField()

	for i := 0; i < numField; i ++ {
		field := t.Field(i)
		fieldValue := ele.Field(i)  // 用指针指向结构体来访问
		if fieldValue.CanSet() {
			fmt.Printf("%s 可以被设置", field.Name)
		}
	}
}


func main() {
	h := Hello()
	//PrintFuncName(*h)
	//PrintFieldName(*h)
	//PrintTypeValue(*h)
	PrintFieldNameFroPtr(h)

}


/**
 * @project prac_grpc
 * @Author 27
 * @Description //TODO
 * @Date 2021/6/6 11:24 6月
 **/
package hello_service

import (
	"fmt"
	"io/ioutil"
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
	// 只能改这个
	FuncField func(name string) (string, error)
}

/*
RPC —— 调用规约
• 参数是一个结构体指针
• 返回值是一个结构体指针和error
• 参数放在 HTTP body 里面进行传输，采 用JSON作为序列化格式
 */



func (h hello) GetOrder(name string) (string, error) {
	return "mock", nil
}

// 改不了它
func (h hello) SayHello(name string) (string, error) {
	return "", nil
	//client := http.Client{}
	//resp, err := client.Get(h.endpoint + name)
	//if err != nil {
	//	//fmt.Printf("%+v", err.Error())
	//	return "", err
	//}
	//
	//data, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	//fmt.Printf("%+v", err)
	//	return "", err
	//}
	//
	//fmt.Printf("%v", data)
	//
	//return string(data), nil
}



// 为指针写的方法
func SetFuncField(val interface{}) {
	v := reflect.ValueOf(val) // 这是指针的反射
	ele := v.Elem()           // 拿到了指针指向的结构体
	t := ele.Type()           // 拿到了指针指向的结构体的类型信息

	numField := t.NumField()

	for i := 0; i < numField; i++ {
		field := t.Field(i)
		fmt.Println(field.Name)
		fieldValue := ele.Field(i) // 用指针指向结构体来访问
		if fieldValue.CanSet() {
			fn := func(args []reflect.Value) (results []reflect.Value) {
				/*
				形式:y := T(x)
				• 如何理解?记住数字类型转换，string 和 []byte 互相转
				• 类似Java强制类型转换
				• 编译器会进行类型检查，不能转换的会编 译错误
				 */
				name := args[0].Interface().(string)
				fmt.Printf("你正在调用方法" + field.Name)
				client := http.Client{}
				resp, err := client.Get("http://localhost:8080/" + name)
				if err != nil {
					fmt.Printf("%+v", err.Error())
					return []reflect.Value{reflect.ValueOf(""), reflect.ValueOf(err)}
				}

				data, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("%+v", err)
					return []reflect.Value{reflect.ValueOf(""), reflect.ValueOf(err)}
				}

				fmt.Printf("%v", data)
				fmt.Println(string(data))
				return []reflect.Value{reflect.ValueOf(string(data)),
					reflect.Zero(reflect.TypeOf(new(error)).Elem())}
				//return []reflect.Value{reflect.ValueOf("Hello, " + name),
				//	reflect.Zero(reflect.TypeOf(new(error)).Elem())}
			}

			fieldValue.Set(reflect.MakeFunc(fieldValue.Type(), fn))
		}
	}

}


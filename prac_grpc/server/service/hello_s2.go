/**
 * @project prac_grpc
 * @Author 27
 * @Description //TODO
 * @Date 2021/6/6 11:24 6月
 **/
package hello_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

func Hello2() *hello2 {
	return &hello2{}
}


type hello2 struct {
	// 只能改这个
	//GetUser func(req *UserReq) (*User, error)
	SayHello2 func(in *Input)(*Output, error)
}


type Input struct {
	Name string
}

type Output struct {
	Msg string
}

/*
RPC —— 调用规约
• 参数是一个结构体指针
• 返回值是一个结构体指针和error
• 参数放在 HTTP body 里面进行传输，采 用JSON作为序列化格式
*/


// 为指针写的方法
func SetFuncField2(val interface{}) {
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

				in := args[0].Interface()
				out := reflect.New(field.Type.Out(0).Elem()).Interface()
				inData, err := json.Marshal(in)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}
				client := http.Client{}

				resp, err := client.Post(
					"http://localhost:8080/",
					"application/json",
					bytes.NewReader(inData))

				if err != nil {
					fmt.Printf("%+v", err.Error())
					return []reflect.Value{reflect.ValueOf(""), reflect.ValueOf(err)}
				}

				data, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(""), reflect.ValueOf(err)}
				}

				err = json.Unmarshal(data, out)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}

				return []reflect.Value{reflect.ValueOf(out),
					reflect.Zero(reflect.TypeOf(new(error)).Elem())}
				//return []reflect.Value{reflect.ValueOf("Hello, " + name),
				//	reflect.Zero(reflect.TypeOf(new(error)).Elem())}
			}

			fieldValue.Set(reflect.MakeFunc(fieldValue.Type(), fn))
		}
	}

}


/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/5/9 18:24 5月
 **/
package main

import (
	"fmt"
	"net/http"
	"strings"
)

func tRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println("form::::", r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme is:", r.URL.Scheme)
	fmt.Println("234:->", r.Form["some_p"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello lalala!") //这个写入到w的是输出到客户端的
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/world", world)

	go func() {
		http.ListenAndServe(":8081", nil)
	}()

	http.ListenAndServe(":8082", nil)
	//http.HandleFunc("/", tRequest)       //设置访问的路由
	//err := http.ListenAndServe(":9090", nil) //设置监听的端口
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
}

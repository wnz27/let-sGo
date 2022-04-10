/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2021-10-29 03:35:12
 * @LastEditTime: 2022-03-18 14:32:15
 * @FilePath: /let-sGo/best_prac_explore/prac_grpc/client/client.go
 * @description: type some description
 */

package main

import (
	"fmt"

	hello_service "fzkprac/best_prac_explore/prac_grpc/server/service"
)

func main() {

	//h := hello_service.Hello()
	// version1
	//msg, err := h.SayHello()
	//if err != nil {
	//	fmt.Printf("%+v", err.Error())
	//	return
	//}
	// version 2
	//msg, _ := h.SayHello("golang")
	//hello_service.SetFuncField(h)
	//msg, _ := h.FuncField("golang")
	h2 := hello_service.Hello2()
	hello_service.SetFuncField2(h2)
	msg, _ := h2.SayHello2(
		&hello_service.Input{Name: "fzkkkkkkk"})

	fmt.Println(msg)

}

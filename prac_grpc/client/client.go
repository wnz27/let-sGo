/**
 * @project prac_grpc
 * @Author 27
 * @Description //TODO
 * @Date 2021/5/23 13:23 5月
 **/
package main

import (
	"fmt"
	hello_service "fzkprac/prac_grpc/server/service"
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

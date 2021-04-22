/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/4/19 22:19 4月
 **/
package main

import (
	"fmt"
	"fzkprac/go_learn/homework/week2_hw"
)

func main() {
	// 向上抛
	res1, err1 := week2_hw.FindSomethingRaise()
	fmt.Println(res1, err1)

	fmt.Println(" =====================   ", " =====================   ")

	// 自己能hold住
	res2, err2 := week2_hw.FindSomethingCanHandle()
	fmt.Println(res2, err2)

	fmt.Println(" =====================   ", " =====================   ")

	// 终止并打日志
	_, _ = week2_hw.ForceFindWithException("test_homework")
	//fmt.Println(res3, err3)

}



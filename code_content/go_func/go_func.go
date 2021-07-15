/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/4/19 21:31 4月
 **/
package main

import (
	"fmt"
	"reflect"
)

// return bigger value
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 函数不定参
func sum_valus(values ...int) {
	fmt.Print(values, " ")
	total := 0
	for _, val := range values {  // 要值不要索引
		total += val
	}
	fmt.Println(total)
}


func hello() []string {
	return nil
}


func main() {
	h := hello
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
		fmt.Println(h, reflect.TypeOf(h))
	}
	// fmt.Println(max(4, 5))

	//sum_valus(1, 2, 3)
	//sum_valus(5, 6)

	// 传数组, 注意语法
	//nums := []int{1, 2, 3, 4, 5}
	//sum_valus(nums...)

	fmt.Println("========================================================================")




}

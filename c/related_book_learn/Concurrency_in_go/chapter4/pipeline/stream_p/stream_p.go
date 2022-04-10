/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/17 01:39 8月
 **/
package main

import "fmt"

func main() {
	multiply := func(value int, multiplier int) int {
		return value * multiplier
	}
	add := func(value, additive int) int {
		return value + additive
	}
	ints := []int{1, 2, 3, 4}
	for _, v := range ints {
		fmt.Println(multiply(add(multiply(v, 2), 1), 2))
	}
}

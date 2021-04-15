/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/4/15 7:33 下午
* Description:
 */
package main

import "fmt"

func main() {
	a := new(struct{})
	b := new(struct{})
	println(a, b, a == b)

	c := new(struct{})
	d := new(struct{})
	fmt.Println(c, d)
	println(c, d, c == d)
}

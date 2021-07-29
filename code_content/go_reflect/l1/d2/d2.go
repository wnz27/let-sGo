/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/7/29 2:35 下午
* Description:
 */
package main

import (
	"fmt"
	"reflect"
)

type Cat struct {
	Name string
}

type MyInt int

func main() {

	var f float64 = 3.5
	t1 := reflect.TypeOf(f)
	fmt.Println(t1.String())

	c := Cat{
		Name: "kitty",
	}
	t2 := reflect.TypeOf(c)
	fmt.Println(t2.String())

	v1 := reflect.ValueOf(f)
	fmt.Println(v1)
	fmt.Println(v1.String())

	v2 := reflect.ValueOf(c)
	fmt.Println(v2)
	fmt.Println(v2.String())


	var i int
	var j MyInt

	i = int(j) // Todo 必须强转

	ti := reflect.TypeOf(i)
	fmt.Println("type of i:", ti.String())

	tj := reflect.TypeOf(j)
	fmt.Println("type of j:", tj.String())

	fmt.Println("kind of i:", ti.Kind())
	fmt.Println("kind of j:", tj.Kind())
}
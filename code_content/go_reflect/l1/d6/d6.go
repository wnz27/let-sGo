/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/7/29 5:23 下午
* Description:
 */
package main

import (
	"fmt"
	"reflect"
)

func Add(a, b int) int {
	return a + b
}

func Greeting(name string) string {
	return "hello " + name
}

func inspectFunc(name string, f interface{}) {
	t := reflect.TypeOf(f)
	fmt.Println(name, "input:")
	for i := 0; i < t.NumIn(); i++ {
		t := t.In(i)
		fmt.Print(t.Name())
		fmt.Print(" ")
	}
	fmt.Println()

	fmt.Println("output:")
	for i := 0; i < t.NumOut(); i++ {
		t := t.Out(i)
		fmt.Print(t.Name())
		fmt.Print(" ")
	}
	fmt.Println("\n===========")
}

func main() {
	inspectFunc("Add", Add)
	inspectFunc("Greeting", Greeting)
}

/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/7/28 4:57 下午
* Description:
 */
package main

import "fmt"

type Animal interface {
	Speak()
}

type Cat struct {
}

func (c Cat) Speak() {
	fmt.Println("Meow")
}

type Dog struct {
}

func (d Dog) Speak() {
	fmt.Println("Bark")
}

func main() {
	var a Animal

	a = Cat{}
	a.Speak()
	fmt.Printf("%v\n", &a)

	a = Dog{}
	a.Speak()
	fmt.Printf("%v\n", &a)
}
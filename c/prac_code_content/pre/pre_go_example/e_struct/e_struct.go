/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/13 12:27 下午
* Description:
 */
package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	return &person{
		name: name,
		age: 42,
	}
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("FFF"))

	s := person{name: "Petson", age: 33}
	fmt.Println(s.name)

	sp :=&s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}

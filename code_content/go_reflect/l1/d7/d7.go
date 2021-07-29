/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/7/29 5:36 下午
* Description:
 */
package main

import (
	"fmt"
	"reflect"
)
func inspectMethod(o interface{}) {
	t := reflect.TypeOf(o)

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)

		fmt.Println(m)
	}
}

type User struct {
	Name    string
	Age     int
}

func (u *User) SetName(n string) {
	u.Name = n
}

func (u *User) SetAge(a int) {
	u.Age = a
}

func main() {
	u := User{
		Name:    "dj",
		Age:     18,
	}
	inspectMethod(&u)
}

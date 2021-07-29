/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/7/29 5:05 下午
* Description:
 */
package main

import (
	"fmt"
	"reflect"
)


func inspectMap(m interface{}) {
	v := reflect.ValueOf(m)
	for _, k := range v.MapKeys() {
		field := v.MapIndex(k)

		fmt.Printf("%v => %v\n", k.Interface(), field.Interface())
	}
}

func main() {
	inspectMap(map[uint32]uint32{
		1: 2,
		3: 4,
	})
}

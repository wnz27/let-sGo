/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/7/29 5:12 下午
* Description:
 */
package main

import (
	"fmt"
	"reflect"
)
func inspectSliceArray(sa interface{}) {
	v := reflect.ValueOf(sa)

	fmt.Printf("%c", '[')
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		fmt.Printf("%v ", elem.Interface())
	}
	fmt.Printf("%c\n", ']')
}

func main() {
	inspectSliceArray([]int{1, 2, 3})
	inspectSliceArray([3]int{4, 5, 6})
}

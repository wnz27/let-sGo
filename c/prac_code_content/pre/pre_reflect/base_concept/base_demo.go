/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/14 6:32 下午
* Description:
 */
package main

import (
	"fmt"
	"reflect"
)

/*
- reflect.Type和 reflect.Value
- reflect.Kind
- NumField() 和 Field() 方法
- Int() 和 String()
 */

type order struct {
	ordId      int
	customerId string
}

func (o order) ttt() {
	fmt.Println(o.ordId)
}

func main() {
	ooo := order{123, "456"}
	//fmt.Println(reflect.Value)

	fmt.Println(reflect.TypeOf(ooo).Name())

	for i := 0; i < reflect.ValueOf(ooo).NumField(); i ++ {
		fmt.Println(reflect.TypeOf(ooo).Field(i).Name)
		fT := reflect.TypeOf(ooo)
		fmt.Println(fT.Field(i).Type.Kind().String())
		fmt.Println(reflect.ValueOf(ooo).Field(i).Int())
	}

	fmt.Println(reflect.TypeOf(ooo).Kind())
	fmt.Println(reflect.TypeOf(ooo).Kind().String())

}

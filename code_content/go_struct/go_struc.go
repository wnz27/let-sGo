/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/7/1 4:07 下午
* Description:
 */
package main

import (
	"fmt"
)

type Test struct {
	name string
}

func (t *Test) Point() {
	fmt.Println(t.name)
}

type person struct {
	name string
}

func hello(num ...int) {
	num[0] = 18
}


func main() {

	ts := []Test{
		{"a"},
		{"b"},
		{"c"},
	}

	for _, t := range ts {
		//fmt.Println("-->", reflect.TypeOf(t))
		defer t.Point()
	}


	sn1 := struct {
		age int
		name string
	}{age: 11, name: "qq"}

	sn2 := struct {
		age int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	sm2 := struct {
		age int
		m map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	fmt.Println(sm1, sm2)

	//if sm1 == sm2 {  // 编译会报错
	//	fmt.Println("sm1 == sm2")
	//}

	// 结构体
	// 1 结构体只能比较是否相等，但是不能比较大小。
	// 2 相同类型的结构体才能够进行比较，结构体是否相同不但与属性类型有关，还与属性顺序相关，sn3 与 sn1
	//就是不同的结构体;
	//sn3:= struct {
	//	name string
	//	age  int
	//}{age:11,name:"qq"}
	// 3 如果 struct 的所有成员都可以比较，则该 struct 就可以通过 == 或 != 进行比较是否相等，比较时逐个项进行 比较，如果每一项都相等，则两个结构体才相等，否则不相等;

	// 可比较的:
	//    常⻅的有 bool、数值型、string、指针、数组等
	//  不可比较:
	//     切片、map、函数等是不能比较的。
	p := new(Test)
	p.name = "SSSS"
	fmt.Println("1 ----> ", p.name)
	fmt.Println("2 ----> ", (*p).name)
	//fmt.Println("3 ----> ", (&p).name)
	//fmt.Println("4 ----> ", p->name)

	var m map[person]int
	p11 := person{"mike"}
	fmt.Println(m[p11])

	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])
}





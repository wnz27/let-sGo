/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/4/24 23:41 4月
 **/
package prac

import "fmt"

func FixString() {
	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n", s2)
}

// 函数作为值、类型
/*
在go 中函数也是一种变量，我们可以通过type 定义它，它的类型就是所有拥有相同的参数，相同返回值的一种类型

 */





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
可以说这一类的函数 函数签名相同
 */

type testInt func(int) bool  // 声明了一种函数类型 需要参数int 返回 bool值

func rem(integer int) (remain int) {
	remain = integer % 2
	return
}

func isOdd(integer int) bool {
	if rem(integer) == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if rem(integer) == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func FuncType() {
	slice := []int {1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)    // 函数当做值来传递了
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven)  // 函数当做值来传递了
	fmt.Println("Even elements of slice are: ", even)
}




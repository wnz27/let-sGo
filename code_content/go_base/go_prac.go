/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/4/24 23:40 4月
 **/
package main
import "C"
import "fmt"

type MyInt1 int
type MyInt2 = int

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func main() {
	// 修改字符串
	//prac.FixString()
	// 测试函数类型
	//prac.FuncType()
	//prac.TMethod()

	/*
	第 10 行代码是基于类型 int 创建了新类型 MyInt1，第 11 行代码是创建了 int 的类型别名 MyInt2，
	注意类型别名的 定义时 = 。
	所以，第 29 行代码相当于是将 int 类型的变􏰀赋值给 MyInt1 类型的变􏰀，
	Go 是强类型语言，编译当 然不通过;而 MyInt2 只是 int 的别名，本质上还是 int，可以赋值。
	第 29 行代码的赋值可以使用强制类型转化 var i1 MyInt1 = MyInt1(i).
	 */
	//var i int =0
	//var i1 MyInt1 = i
	//var i2 MyInt2 = i
	//fmt.Println(i1,i2)
	var x interface{} = nil
	fmt.Println(x,y,z,k,p)

}

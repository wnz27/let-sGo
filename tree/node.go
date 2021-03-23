/**
 * @project let-sGo
 * @Author 27
 * @Description struct learn
 * @Date 2021/3/23 23:43 3月
 **/
package main


import (
	"fmt"
)

/*  总结
指针接收者 VS 值接收者
要改变内容必须使用指针接收者
结构过大也使用指针接收者，省空间和性能
一致性：如有指针接收者，最好都是指针接收者

值接收者是go语言特有！
值/指针接收者均可接收值/指针， 一个方法创建者可以随意切换是值接收还是指针接收，调用处可以不做修改
*/


type treeNode struct {
	value int
	left, right *treeNode
}

/*
func (variable_name variable_type) function_name ([parameter list]) [return_types]
{
}
*/
// 函数接收者
func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

// 只有使用指针接收者 才可以改变结构的内容
/*
！！！！！！！  nil 指针也可以调用方法    ！！！！！！！
*/
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node!!!")
		return
	}
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

// 无法修改成功
func (node treeNode) setValue0(value int) {
	node.value = value
}

/*
C++ 局部变量分配在栈上，函数一旦退出，局部变量立刻被销毁，如果要传出去必须要在堆上分配，堆上分配的话要手动释放，这是C++的做法。
Java 几乎所有东西都是分配在堆上，我们都要用new，所以才会有垃圾回收机制
*/
// 结构是分配在堆还是栈上
/*
go 无需关心堆还是栈，有垃圾回收，比如这个工厂函数，可能编译器会优化比如如果这个返回了，那么可能在堆上分配，如果在不返回就在栈上分配，
函数退出就销毁。
*/

// 自定义工厂函数 返回了局部变量的地址
func createNode(value int) * treeNode{
	return &treeNode{value: value}  // 相当于在函数体建了个局部变量给别人用，说明局部变量也可以返回给别人用
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}  // 零值
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode) // 返回的是* 也就是指针，地址
	root.left.right = createNode(2)

	root.right.left.setValue(4)
	root.right.left.print()
	fmt.Println()

	root.right.left.setValue0(55555)
	root.right.left.print()
	fmt.Println()
	// 看函数定义要什么， 函数要指针就是把调用者的地址给函数，如果函数要值，会把调用者的值copy给函数
	root.traverse()

	//pRoot := &root
	//pRoot.print()
	//pRoot.setValue(200)
	//pRoot.print()

	qRoot := treeNode{}  // 都是零值
	var qRoot1 *treeNode  // nil
	fmt.Println(qRoot, qRoot1)

	fmt.Println()

	//qRoot1.setValue(222)
	//qRoot1 = &root
	//qRoot1.setValue(333)
	//qRoot1.print()


	// 切片里可以有一些省略
	//nodes := []treeNode {
	//	{value: 4},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)

}

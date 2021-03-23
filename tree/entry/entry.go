/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/3/24 01:37 3月
 **/
package main

import (
	"fmt"
	"fzkprac/tree"
)

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}  // 零值
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node) // 返回的是* 也就是指针，地址
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()

	root.Right.Left.SetValue0(55555)
	root.Right.Left.Print()
	fmt.Println()
	// 看函数定义要什么， 函数要指针就是把调用者的地址给函数，如果函数要值，会把调用者的值copy给函数
	root.Traverse()

	//pRoot := &root
	//pRoot.print()
	//pRoot.setValue(200)
	//pRoot.print()

	qRoot := tree.Node{}  // 都是零值
	var qRoot1 *tree.Node  // nil
	fmt.Println(qRoot, qRoot1)

	fmt.Println()

	//qRoot1.setValue(222)
	//qRoot1 = &root
	//qRoot1.setValue(333)
	//qRoot1.print()


	// 切片里可以有一些省略
	//tree.Nodes := []tree.Node {
	//	{Value: 4},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(tree.Nodes)

}


/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-09-20 08:37:35
 * @LastEditTime: 2022-09-20 08:49:15
 * @FilePath: /let-sGo/c/al_practice/ds/tree_p.go
 * @description: type some description
 */

package main

import "fmt"

//Tree struct
type Tree struct {
	LeftNode  *Tree
	Value     int
	RightNode *Tree
}

func (tree *Tree) insert(m int) {
	if tree != nil {
		if tree.LeftNode == nil {
			tree.LeftNode = &Tree{nil, m, nil}
		} else {
			if tree.RightNode == nil {
				tree.RightNode = &Tree{nil, m, nil}
			} else {
				if tree.LeftNode != nil {
					tree.LeftNode.insert(m)
				} else {
					tree.RightNode.insert(m)
				}
			}
		}
	} else {
		tree = &Tree{nil, m, nil}
	}
}

func print(tree *Tree) {
	if tree != nil {
		fmt.Println(" Value", tree.Value)
		fmt.Printf("Tree Node Left")
		print(tree.LeftNode)
		fmt.Printf("Tree Node Right")
		print(tree.RightNode)
	} else {
		fmt.Printf("Nil\n")
	}
}

// func main() {
// 	var tree *Tree = &Tree{nil, 1, nil}
// 	print(tree)
// 	tree.insert(3)
// 	print(tree)
// 	tree.insert(5)
// 	print(tree)
// 	tree.LeftNode.insert(7)
// 	print(tree)
// }

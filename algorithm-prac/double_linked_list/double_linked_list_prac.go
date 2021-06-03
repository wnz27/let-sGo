/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/6/3 5:23 下午
* Description:
 */
package main

import "fmt"

type Node struct {
	Value    int
	Previous *Node
	Next     *Node
}

var root = new(Node)

func addNode(t *Node, v int) int {
	if root == nil {
		t = &Node{v, nil, nil}
		root = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node is exist:", v)
		return -1
	}

	if t.Next == nil {
		temp := t
		t.Next = &Node{v, temp, nil}
		return -2
	}
	return addNode(t.Next, v)
}


func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for t != nil {
		fmt.Println()
	}
}

func main() {

}

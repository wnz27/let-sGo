/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/13 1:53 下午
* Description:
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}

func traverse(t * Tree)  {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Println(t.Value, " ")
	traverse(t.Right)
}

// 向树中填充随机数
func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2 * n; i ++ {
		temp := rand.Intn(n * 2)
		t = insert(t, temp)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}

	if v == t.Value {
		return t
	}

	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}

	t.Right = insert(t.Right, v)
	return t
}


func main()  {
	tree := create(10)
	traverse(tree)
	fmt.Println()
	tree = insert(tree, -10)
	tree = insert(tree, -2)
	traverse(tree)
	fmt.Println()
	fmt.Println("The value of the root of the tree is", tree.Value)
	fmt.Println("The value of the root of the tree is", tree.Value)
}






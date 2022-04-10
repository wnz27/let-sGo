/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/6/3 3:02 下午
* Description:
 */
package main

import "fmt"

const SIZE = 15

type Node struct {
	Value int
	Next *Node
}

type HashTable struct {
	Table map[int]*Node
	Size int
}


func hashFunction(i, size int) int {
	return i % size
}

func insert(hash *HashTable, value int) int {
	key := hashFunction(value, hash.Size)
	element := Node{Value: value, Next: hash.Table[key]}
	hash.Table[key] = &element
	return 0
}

func traverse(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
		}
	}
	fmt.Println()
}

func lookup(hash *HashTable, value int) bool {
	key := hashFunction(value, hash.Size)
	if hash.Table[key] != nil {
		t := hash.Table[key]
		for t != nil {
			if t.Value == value {
				return true
			}
			t = t.Next
		}
	}
	return false
}

func main () {
	table := make(map[int]*Node, SIZE)
	hash := &HashTable{Table: table, Size: SIZE}
	fmt.Println("Numbder of spaces:", hash.Size)
	for i := 0; i < 120; i++ {
		insert(hash, i)
	}
	traverse(hash)
}


/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/3 6:30 下午
* Description:
 */
package main

import (
	"sync"
	"time"
)

var a string
var done bool
var once sync.Once

func setup() {
	a = "hello, world"
	done = true
}

func doprint() {
	if !done {
		once.Do(setup)
	}
	print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}

func main(){
	twoprint()
	time.Sleep(time.Millisecond * 500)
}

// todo： 但是自己模拟没有模拟出来, 可能编译乱序的时候会出现
// but there is no guarantee that, in doprint, observing the write to done implies observing the write to a.
// This version can (incorrectly) print an empty string instead of "hello, world".
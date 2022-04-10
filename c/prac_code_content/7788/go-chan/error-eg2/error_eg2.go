/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/3 6:31 下午
* Description:
 */
package main

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func main() {
	go setup()
	for !done {
	}
	print(a)
}
// todo  自己也没模拟出失败， 可能编译乱序的时候会出现
// As before, there is no guarantee that, in main, observing the write to done implies observing the write to a,
// so this program could print an empty string too.
// Worse, there is no guarantee that the write to done will ever be observed by main,
// since there are no synchronization events between the two threads.
// The loop in main is not guaranteed to finish.


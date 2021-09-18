/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/3 9:20 下午
* Description:
 */
package main

type T struct {
	msg string
}

var g *T

func setup() {
	t := new(T)
	t.msg = "hello, world"
	g = t
}

func main() {
	go setup()
	for g == nil {
	}
	print(g.msg)
}

// todo 没有复现不保证的情况， 可能编译乱序的时候会出现
// Even if main observes g != nil and exits its loop,
//there is no guarantee that it will observe the initialized value for g.msg.
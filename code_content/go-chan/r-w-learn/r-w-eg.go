/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/3 6:14 下午
* Description:
 */
package main

var a, b int

func w() {
	a = 1
	b = 2
}

func r() {
	print(b)
	print(a)
}

func main() {
	go w()
	//time.Sleep(time.Millisecond*50)  // 给协程时间写入它就会有值，但这种行为是不受我们控制的
	r()
}

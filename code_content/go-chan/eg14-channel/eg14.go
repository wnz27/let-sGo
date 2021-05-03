/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/3 5:22 下午
* Description:
 */
package main

// buffer or un buffered


import (
	"fmt"
	"time"
)

var c = make(chan int)
var a string

func f() {
	a = "hello, world"
	time.Sleep(time.Second*1)
	<-c
}

func main() {
	go f()
	c <- 0
	fmt.Println("123")
	print(a)
}
// If the channel were buffered (e.g., c = make(chan int, 1)) then the program would not be guaranteed to print "hello, world".
//(It might print the empty string, crash, or do something else.)

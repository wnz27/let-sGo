/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/13 4:08 下午
* Description:
 */
package main

import "fmt"

func ping(pings chan <- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <-msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

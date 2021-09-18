/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/13 3:53 下午
* Description:
 */
package main

import "fmt"

func main() {
	messages := make(chan string)
	defer close(messages)

	go func() {
		messages <- "ping"
	}()

	msg := <- messages
	fmt.Println(msg)
}



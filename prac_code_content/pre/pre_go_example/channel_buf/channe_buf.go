/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/13 3:55 下午
* Description:
 */
package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	defer close(messages)
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

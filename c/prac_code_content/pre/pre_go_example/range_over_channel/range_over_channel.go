/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/13 5:42 下午
* Description:
 */
package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}

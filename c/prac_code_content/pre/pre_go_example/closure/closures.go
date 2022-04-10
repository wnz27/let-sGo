/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/13 11:54 上午
* Description:
 */
package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i ++
		return i
	}
}

func main() {
	nextI := intSeq()

	fmt.Println(nextI())
	fmt.Println(nextI())
	fmt.Println(nextI())

	newNextI := intSeq()
	fmt.Println(newNextI())
}

/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/13 3:34 下午
* Description:
 */
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i ++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

}

/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/13 3:59 下午
* Description:
 */
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working ...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<- done
}

/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/13 01:49 8月
 **/
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	go func() {
		defer close(done)
		time.Sleep(5 * time.Second)
	}()

	workCounter := 0
	loop:
	for {
		select {
		case <- done:
			break loop
		default:
		}
		// 模拟工作行为
		fmt.Printf("work: %v\n", workCounter)
		workCounter ++
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}

/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/13 01:41 8月
 **/
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var c1, c2 <- chan int
	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}
}

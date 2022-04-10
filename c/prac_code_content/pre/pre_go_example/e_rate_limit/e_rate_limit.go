/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/15 8:49 下午
* Description:
 */
package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 0; i < 5; i ++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i ++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i ++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	st1 := time.Now()
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
	fmt.Println(time.Since(st1))
}

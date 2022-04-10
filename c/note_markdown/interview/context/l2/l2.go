/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/7/20 4:11 下午
* Description:
 */
package main

import "context"

// 结合goroutine 常见以下形式
func l2(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

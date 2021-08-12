/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/12 01:29 8月
 **/
package main

import (
	"fmt"
	"sync"
)

func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i ++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<- begin  // goroutine会一直等待，直到它被告知可以继续
			fmt.Printf("%v has begun \n", i)
		}(i)
	}
	fmt.Println("Unblocking goroutines ......")
	close(begin)  // 关闭channel 从而同时打开所有的goroutine
	wg.Wait()
}

/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/15 03:10 8月
 **/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")  // 在goroutine 成功终止时打印出一条消息
			defer close(randStream)
			for {
				randStream <- rand.Int()
			}
		}()
		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i ++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}

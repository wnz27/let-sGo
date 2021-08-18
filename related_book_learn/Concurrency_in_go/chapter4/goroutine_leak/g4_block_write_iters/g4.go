/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/15 03:55 8月
 **/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)

	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)

	// 模拟正在进行的工作
	time.Sleep(1 * time.Second)
	fmt.Println("do some work use 1s")

}

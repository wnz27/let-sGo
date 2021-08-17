/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/18 01:15 8月
 **/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i ++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:  // 这里书上错了！！！！！
				}
			}
		}()
		return takeStream
	}

	repeatFn := func(
		done <-chan interface{},
		fn func() interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}
	done := make(chan interface{})
	defer close(done)

	rand := func() interface{} {return rand.Int()}

	for num := range take(done, repeatFn(done, rand), 10) {
		fmt.Println(num)
	}
}

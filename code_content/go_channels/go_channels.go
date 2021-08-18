/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/4/19 20:11 4月
 **/
package main

import "fmt"

/*
func c1() {
	c := make(chan int)
	go func() {
		list.Sort()
		c <- 1
	}()
	<-c
}
*/

func main() {
	//intStream := make(chan interface{})
	//
	//go func() {
	//	defer close(intStream)
	//	for i := 0; i < 3; i++ {
	//		intStream <- i
	//	}
	//}()
	//for i := range intStream {
	//	fmt.Println(i)
	//}

	repeat := func(
		done <-chan interface{},
		values ...interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}()
		return valueStream
	}
	done2 := make(chan interface{})
	defer close(done2)
	for nnn := range repeat(done2, 2, 3, 4) {
		fmt.Println(nnn)
	}
}



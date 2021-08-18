/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/14 01:18 8月
 **/
package main

import "fmt"

func main() {
	data := make([]int, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}

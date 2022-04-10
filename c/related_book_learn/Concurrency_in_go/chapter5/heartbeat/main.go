/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/9/3 01:43 9月
 **/
package main

import "fmt"

func main() {
	rr := make(chan int)
	close(rr)
	for i := range rr {
		fmt.Println(i)
	}
}



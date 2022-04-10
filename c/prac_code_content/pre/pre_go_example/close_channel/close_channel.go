/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/13 5:18 下午
* Description:
 */
package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j ++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)

	fmt.Println("sent all jobs")

	<-done
}

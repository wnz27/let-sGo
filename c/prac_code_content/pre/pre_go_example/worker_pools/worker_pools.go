/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/15 4:36 下午
* Description:
 */
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started jobs", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished jobs", j)
		results <- j * 2
	}
}

func main() {
	startTime := time.Now()

	const numJobs = 9
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w ++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j ++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}

	duration := time.Since(startTime)
	fmt.Println("real: ", duration)
}

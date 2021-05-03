/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/3 5:32 下午
* Description:
 */
package main

var limit = make(chan int, 3)

func main() {
	limit <- 1
	<-limit
	limit <- 1
	limit <- 1
	<-limit
	<-limit
}

/*
This program starts a goroutine for every entry in the work list,
but the goroutines coordinate using the limit channel to ensure that at most three are running work functions at a time.
var limit = make(chan int, 3)

func main() {
	for _, w := range work {
		go func(w func()) {
			limit <- 1
			w()
			<-limit
		}(w)
	}
	select{}
}
 */

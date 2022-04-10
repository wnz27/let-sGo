/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/15 02:03 8月
 **/
package main

import "fmt"

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				// 做些有趣的事
				fmt.Println(s)
			}
		}()
		return completed
	}

	res := doWork(nil)
	// 也许有其他的操作需要进行
	<-res
	fmt.Println("Done")
}

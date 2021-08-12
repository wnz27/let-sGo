/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/13 01:28 8月
 **/
package main

import (
	"fmt"
	"time"
)

func main() {
	var c <-chan int
	select {
	case <-c:  // 这个case永远不会被执行，因为我们是从 nil channel 读取的。死锁（永久阻塞）
	case <- time.After(1 * time.Second):
		fmt.Println("Time out!")
	//default:  // 打开注释，则不会走timeout了
	//	fmt.Println("AAAAAAAA")
	}
}

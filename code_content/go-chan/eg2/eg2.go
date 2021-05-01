/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/5/1 10:51 5月
 **/
package main

import (
	"fmt"
	"time"
)

// 使用for 从通道接收数据
func main(){
	ch := make(chan int)

	go func() {
		for i := 3; i >= 0 ; i-- {
			ch <- i

			// 每次发送完时等待
			time.Sleep(time.Second)
		}
	}()

	// 遍历接收通道数据
	for data := range ch {
		// 打印通道数据
		fmt.Println(data)

		// 当遇到数据0时，退出
		if data == 0 {
			break
		}
	}
}

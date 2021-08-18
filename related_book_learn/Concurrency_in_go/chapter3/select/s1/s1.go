/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/13 00:49 8月
 **/
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5*time.Second)  // 等待5s后关闭channel
		close(c)
	}()

	fmt.Println("Bocking on read ...")
	select {
	//case <-c:
	//	fmt.Println("123")
	case <-c:  // 尝试在channel上读取数据。 注意，在编写这段代码时，我们不需要select语句。可以简单地使用 <-c，但是我们在这个示例中进行扩展
		fmt.Printf("Unblock %v later.\n", time.Since(start))
	}
}

/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/15 02:08 8月
 **/
package main

import "fmt"

func main() {
	var nilChan chan interface{}
	// 测试 nil chan 读值 deadlock
	//for s := range nilChan {
	//	fmt.Println(s)
	//}
	// 测试关闭读值  会panic
	close(nilChan)
	a := <-nilChan
	fmt.Println(a)
}

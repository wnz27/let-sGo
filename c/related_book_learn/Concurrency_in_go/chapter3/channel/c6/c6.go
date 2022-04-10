/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/12 23:55 8月
 **/
package main

import (
	"fmt"
)

func main() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)  // 实例化一个缓冲channel。因为知道将产生6个结果，我们创建一个有5个缓冲的channel，这样goroutine就能尽快完成
		go func() {  // 启动一个匿名的goroutine，它在resultStream上执行写操作。注意，我们已经颠倒了如何创建goroutine。它现在被封装在外围函数中
			defer close(resultStream)  // 确保一旦执行完成resultStream就会关闭。作为channel所有者，这是我们必须做的。
			for i := 0; i <=5; i ++ {
				resultStream <- i
				//if i == 5{
				//	time.Sleep(2*time.Second)
				//	resultStream <- i
				//} else {
				//	resultStream <- i
				//}
			}
		}()
		return resultStream  // 在这里我么你返回channel。由于返回值被声明为一个只读channel，因此resultStream将隐式的转换为只读消费者
	}

	resultStream := chanOwner()
	// chan不关闭这个循环不会退出
	for result := range resultStream {  // 遍历resultStream，作为消费者，我们只关心阻塞和channel的关闭
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}


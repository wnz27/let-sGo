/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/8 19:54 8月
 **/
package main

import (
	"fmt"
	"sync"
	"time"
)

/*
扩展c1 的例子，并显示等式的两边：等待信号的goroutine和发送信号的goroutine。
假设我们有一个固定长度为2的队列，还有10条我们想要推送到队列中的数据。
我们想要在有容量的情况下，尽快进入队列，所以就希望在队列中有空间时能立即得到通知。
让我们尝试使用Cond来管理这种调度：
 */
func main() {
	// 使用sync.Mutex 作为锁
	c := sync.NewCond(&sync.Mutex{})
	// 创建一个长度为0的切片，因为我们最终会添加10条数据
	// 所以用10的容量实例化它
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		// 再次进入临界区，以便我们可以修改与条件相关的数据。
		c.L.Lock()

		// 通过将切片的头部重新分配给第二条数据来模拟对一条数据的排队
		queue = queue[1:]
		fmt.Println("Remove from queue")
		// 退出条件的临界区， 因为成功删除一条数据
		c.L.Unlock()

		// 我们让一个等待的goroutine知道发生了什么
		c.Signal()
	}

	for i := 0; i < 10; i ++ {
		// 我们通过在条件的锁存器上调用锁来进入临界区。
		c.L.Lock()

		// 检查一个循环队列的长度，这很重要，因为在这种情况下的信号并不一定
		// 意味着是你所等的信号，也可能只是发生了什么
		for len(queue) == 2 {
			// 调用Wait 这将暂停main goroutine 直到一个信号的条件已经发送
			c.Wait()
		}

		fmt.Printf("%d Adding to queue\n", i)
		queue = append(queue, struct {}{})

		// 创建了一个新的goroutine， 它将在一秒钟后删除一个元素
		go removeFromQueue(1 * time.Second)

		// 退出条件的临界区，因为我们已成功读取了一条数据
		c.L.Unlock()
	}

	/*
	输出：
	Adding to queue
	Adding to queue
	Remove from queue
	Adding to queue
	Remove from queue
	Adding to queue
	Remove from queue
	Remove from queue
	Adding to queue
	Adding to queue
	Remove from queue
	Adding to queue
	Remove from queue
	Adding to queue
	Remove from queue
	Adding to queue
	Remove from queue
	Adding to queue
	 */

	/*
	该程序，成功的将所有的10条数据添加到队列中（并且在它有机会将前两条删除之前退出）。
	它也总是等待，知道至少有一条数据被写入队列，然后再执行另一条数据。
	 */

	/*
	这个例子中，我们还有一个新方法：Signal。
	这是Cond类型提供的两种方法中的一种，它提供通知goroutine阻塞的调用Wait，条件已经被触发。

	另一种方法叫Broadcast。运行时内部维护一个FIFO列表，等待接收信号；
	Signal 发现等待最长时间的goroutine 并通知他，而broadcast 向所有等待的goroutine 发送信号。
	Broadcast 可以说是这两种方法中比较有趣的一种，因为它提供了一种同时与多个goroutine通信的方法。

	我们可以通过channel 对信号进行简单的复制，但是重复调用Broadcast的行为将会更加困难。
	此外，与利用channel相比，Cond类型的性能要高很多。
	 */

}

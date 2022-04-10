/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/8 17:42 8月
 **/
package main

import (
	"sync"
	"testing"
)

/*
来看看在OS 线程和goroutine 之间切换的上下文的相对性能。首先，我们将利用Linux 的内置基准测试套件来度量在
相同核心的两个线程之间发送消息需要多长时间：
taskset -c 0 perf bench sched pipe -T
运行结果（个人华为云服务器centos 7.7 执行结果） 需要安装perf
# Running 'sched/pipe' benchmark:
# Executed 1000000 pipe operations between two threads

     Total time: 2.124 [sec]

       2.124964 usecs/op
         470596 ops/sec

这个基准是度量了在线程上发送和接收消息所需的时间，因此我们将计算结果并将其除以2，用了1.05 微秒来进行上下文切换，
看起来不算糟糕。

下面我们来检查下goroutine的上下文切换
 */
func BenchmarkContextSwitch(b *testing.B) {
	// 下面示例将创建两个goroutine 并在它们之间发送一条消息：
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i<b.N; i ++ {
			c <- token
		}
	}

	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i ++ {
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
	/*
	运行结果
	BenchmarkContextSwitch   6978378               171.1 ns/op
	PASS
	ok      fzkprac/related_book_learn/Concurrency_in_go/chapter3/goroutine/g3      1.850s
	*/

	/*
	相比与操作系统的线程，1.05 微秒 这个是0.171微秒, 提升了快百分百了。
	所以很难判断多少goroutine 会导致上下文切换过于频繁，但是上限很可能不会成为使用goroutine的障碍

	创建goroutine 非常廉价，只有当你已经证明了它们是性能问题的根本原因后，你才应该讨论它们的成本
	 */

}

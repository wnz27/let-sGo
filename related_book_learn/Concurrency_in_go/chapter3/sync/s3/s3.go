/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/8 19:10 8月
 **/
package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

/*
进入和退出临界区是由消耗的，所以一般人会尽量减少在临界区的时间。
这样做的一个策略是减少临界区的范围。
可能存在需要在多个并发进程之间共享内存的情况，但可能这些进程不是都需要读写此内存。
如果是这样，你可以利用不同类型的互斥对象：sync.RWMutex

sync.RWMutex 在概念上和互斥是一样的：它保护着对内存的访问，然而，RWMutex让你对内存有了更多的控制。
你可以请求一个锁用于读处理，在这种情况下，你将被授予访问权限，除非该锁被用于写处理。
这意味这，任意数量的读消费者可以持有一个读锁，只要没有其他事物持有一个写锁。
 */
func main() {
	// 这个例子，演示了一个生产者，它不像代码中创建的众多消费者那样活跃:

	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i := 5; i > 0; i -- {
			l.Lock()
			l.Unlock()
			time.Sleep(1)
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()

		go producer(&wg, mutex)
		for i := count ; i > 0; i-- {
			go observer(&wg, rwMutex)
		}
		wg.Wait()

		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")
	for i := 0; i < 20; i ++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}
	/* 我的输出
	Readers  RWMutex       Mutex
	1        37.474µs      2.775µs
	2        4.442µs       2.425µs
	4        6.169µs       7.946µs
	8        53.48µs       43.488µs
	16       51.506µs      48.694µs
	32       16.807µs      28.747µs
	64       209.555µs     184.166µs
	128      125.987µs     98.923µs
	256      145.718µs     125.084µs
	512      253.725µs     258.765µs
	1024     404.564µs     354.162µs
	2048     611.634µs     677.153µs
	4096     1.239702ms    1.375547ms
	8192     2.523513ms    2.498725ms
	16384    5.069722ms    5.131715ms
	32768    10.518789ms   10.530901ms
	65536    20.813154ms   20.411729ms
	131072   41.991906ms   40.036437ms
	262144   86.023476ms   80.420633ms
	524288   173.939331ms  170.703012ms
	*/

	/*
	todo （不是非常理解）
	 这个特殊的例子中减少了临界区的大小，
	 实际上只给开始2^13个读消费者的返回信息。
	 这将取决于你的临界区在做什么，但是通常建议使用RWMutex，而不是Mutex，因为它在逻辑上更加合理。
	 */

}

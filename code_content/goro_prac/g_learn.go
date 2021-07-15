/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/6/23 11:37 下午
* Description:
 */

/*
实现控制并发的方式，大致可分成以下三类：

全局共享变量

channel通信

Context包
 */

/*
todo 1、这是最简单的实现控制并发的方式，实现步骤是：
	- 声明一个全局变量；
	- 所有子goroutine共享这个变量，并不断轮询这个变量检查是否有更新；
	- 在主进程中变更该全局变量；
	- 子goroutine检测到全局变量更新，执行相应的逻辑。
*/
package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"context"
)

func type1() {
	running := true

	f := func() {
		for running {
			fmt.Println("sub proc running...")
			time.Sleep(1 * time.Second)
		}
		fmt.Println("sub proc exit")
	}

	go f()
	go f()
	go f()

	time.Sleep(2 * time.Second)

	running = false

	time.Sleep(3 * time.Second)

	fmt.Println("main proc exit")
}
/*
todo 全局变量的优势是简单方便，不需要过多繁杂的操作，通过一个变量就可以控制所有子goroutine的开始和结束；
 缺点是功能有限，由于架构所致，该全局变量只能是多读一写，否则会出现数据同步问题，当然也可以通过给全局变量加锁来解决这个问题，
 但那就增加了复杂度，另外这种方式不适合用于子goroutine间的通信，因为全局变量可以传递的信息很小；还有就是主进程无法等待所有子goroutine退出，
 因为这种方式只能是单向通知，所以这种方法只适用于非常简单的逻辑且并发量不太大的场景，一旦逻辑稍微复杂一点，这种方法就有点捉襟见肘。

*/


/*
TODO channel通信
 Channel是Go中的一个核心类型，你可以把它看成一个管道，通过它并发核心单元就可以发送或者接收数据进行通讯(communication)。
	要想理解 channel 要先知道 CSP 模型：
	CSP 是 Communicating Sequential Process 的简称，中文可以叫做通信顺序进程，是一种并发编程模型，由 Tony Hoare 于 1977 年提出。
	简单来说，CSP 模型由并发执行的实体（线程或者进程）所组成，实体之间通过发送消息进行通信，这里发送消息时使用的就是通道，或者叫 channel。
	CSP 模型的关键是关注 channel，而不关注发送消息的实体。Go 语言实现了 CSP 部分理论，goroutine 对应 CSP 中并发执行的实体，channel 也就对应着 CSP 中的 channel。
	也就是说，CSP 描述这样一种并发模型：多个Process 使用一个 Channel 进行通信, 这个 Channel 连结的 Process 通常是匿名的，消息传递通常是同步的（有别于 Actor Model）。

 */
func consumer1(stop <-chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("exit sub goroutine")
			return
		default:
			fmt.Println("running...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
func type2() {
	stop := make(chan bool)
	var wg sync.WaitGroup
	// Spawn example consumers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(stop <-chan bool) {
			defer wg.Done()
			consumer1(stop)
		}(stop)
	}
	waitForSignal()
	close(stop)
	fmt.Println("stop all jobs")
	wg.Wait()
}

func waitForSignal() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)
	signal.Notify(sigs, syscall.SIGTERM)
	<- sigs
}

/*
首先了解下channel，可以理解为管道，它的主要功能点是：

队列存储数据

阻塞和唤醒goroutine

channel 实现集中在文件 runtime/chan.go 中，channel底层数据结构是这样的：

 1 type hchan struct {
 2    qcount   uint           // 队列中数据个数
 3    dataqsiz uint           // channel 大小
 4    buf      unsafe.Pointer // 存放数据的环形数组
 5    elemsize uint16         // channel 中数据类型的大小
 6    closed   uint32         // 表示 channel 是否关闭
 7    elemtype *_type // 元素数据类型
 8    sendx    uint   // send 的数组索引
 9    recvx    uint   // recv 的数组索引
10    recvq    waitq  // 由 recv 行为（也就是 <-ch）阻塞在 channel 上的 goroutine 队列
11    sendq    waitq  // 由 send 行为 (也就是 ch<-) 阻塞在 channel 上的 goroutine 队列
12
13    // lock protects all fields in hchan, as well as several
14    // fields in sudogs blocked on this channel.
15    //
16    // Do not change another G's status while holding this lock
17    // (in particular, do not ready a G), as this can deadlock
18    // with stack shrinking.
19    lock mutex
20}

从源码可以看出它其实就是一个队列加一个锁（轻量），代码本身不复杂，但涉及到上下文很多细节，故而不易通读，
有兴趣的同学可以去看一下，我的建议是，从上面总结的两个功能点出发，一个是 ring buffer，用于存数据；
一个是存放操作（读写）该channel的goroutine 的队列。

buf是一个通用指针，用于存储数据，看源码时重点关注对这个变量的读写

recvq 是读操作阻塞在 channel 的 goroutine 列表，sendq 是写操作阻塞在 channel 的 goroutine 列表。
列表的实现是 sudog，其实就是一个对 g 的结构的封装，看源码时重点关注，是怎样通过这两个变量阻塞和唤醒goroutine的

 */


/*
todo Context通常被译作上下文，它是一个比较抽象的概念。在讨论链式调用技术时也经常会提到上下文。
 一般理解为程序单元的一个运行状态、现场、快照，而翻译中上下又很好地诠释了其本质，上下则是存在上下层的传递，上会把内容传递给下。
 在Go语言中，程序单元也就指的是Goroutine。
 每个Goroutine在执行之前，都要先知道程序当前的执行状态，通常将这些执行状态封装在一个Context变量中，传递给要执行的Goroutine中。
 上下文则几乎已经成为传递与请求同生存周期变量的标准方法。在网络编程下，当接收到一个网络请求Request，在处理这个Request的goroutine中，
 可能需要在当前gorutine继续开启多个新的Goroutine来获取数据与逻辑处理（例如访问数据库、RPC服务等），
 即一个请求Request，会需要多个Goroutine中处理。而这些Goroutine可能需要共享Request的一些信息；
 同时当Request被取消或者超时的时候，所有从这个Request创建的所有Goroutine也应该被结束。
 */

type favContextKey string

func type3() {
	wg := &sync.WaitGroup{}
	values := []string{"https://www.baidu.com/", "https://www.zhihu.com/"}
	ctx, cancel := context.WithCancel(context.Background())

	for _, url := range values {
		wg.Add(1)
		subCtx := context.WithValue(ctx, favContextKey("url"), url)
		go reqURL(subCtx, wg)
	}

	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()

	wg.Wait()
	fmt.Println("exit main goroutine")
}

func reqURL(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	url, _ := ctx.Value(favContextKey("url")).(string)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("stop getting url:%s\n", url)
			return
		default:
			r, err := http.Get(url)
			if r.StatusCode == http.StatusOK && err == nil {
				body, _ := ioutil.ReadAll(r.Body)
				subCtx := context.WithValue(ctx, favContextKey("resp"), fmt.Sprintf("%s%x", url, md5.Sum(body)))
				wg.Add(1)
				go showResp(subCtx, wg)
			}
			r.Body.Close()
			//启动子goroutine是为了不阻塞当前goroutine，这里在实际场景中可以去执行其他逻辑，这里为了方便直接sleep一秒
			// doSometing()
			time.Sleep(time.Second * 1)
		}
	}
}

func showResp(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop showing resp")
			return
		default:
			//子goroutine里一般会处理一些IO任务，如读写数据库或者rpc调用，这里为了方便直接把数据打印
			fmt.Println("printing ", ctx.Value(favContextKey("resp")))
			time.Sleep(time.Second * 1)
		}
	}
}

func main() {
	type3()
}



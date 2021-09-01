# 心跳
心跳是并发进程向外界发出信号的一种方式。这个说法来自人体解剖学，在解剖学中心跳
反映了观察者的生命体征。心跳在Go语言之前就已经存在，而且一直非常有效。

在设计并发程序时，一定要考虑到超时和取消。如果从一开始就忽略超时和取消，
然后在后期尝试加入它们，这有点像在蛋糕烤好后再加鸡蛋。

在并发编程中，有几个原因使心跳变得格外有用。它允许我们对系统有深入的了解，当系统
工作不正常时，它可以对系统进行测试。

本节将讨论两种不同类型的心跳：
- 在一段时间间隔内发出的心跳。
- 在工作单元开始时发出的心跳。

在一段时间间隔上发出的心跳对并发代码很有用，尤其是当它在处于等待状态。因为你不知道
新的事件什么时候会被触发，你的goroutine 可能会在等待某件事情发生的时候挂起。
心跳是告诉监听程序一切安好的一种方式，而静默状态也是预料之中的。

下面代码演示了一个会发出心跳的 goroutine:
```go
doWork := func(
		done <-chan interface{},
		pulseInterval time.Duration,
	) (<-chan interface{}, <-chan time.Time) {
		// 我们建立了一个发送心跳的 channel 我们把这个返回给 doWork
		heartbeat := make(chan interface{})
		results := make(chan time.Time)
		go func() {
			defer close(heartbeat)
			defer close(results)

			// 我们设定心跳的间隔时间为我们接到的 pulseInterval。每隔一个
			// pulseInterval 的时长都会有一些消费者读取这个channel
			pulse := time.Tick(pulseInterval)
			// 这是另一个用来模拟心跳行为的channel。我们选择的持续时间大于 pulseInterval,
			// 这样我们就能看到从 goroutine 中发出的一些心跳。
			workGen := time.Tick(2 * pulseInterval)

			sendPulse := func() {
				select {
				case heartbeat <-struct{}{}:
				default: // 这里我们加入一个默认语句。我们必须时刻警惕这样一个事实：
				// 可能会没有人接收我们的心跳。从goroutine 发出的信息是重要的，但心跳却不一定重要。
				}
			}

			sendResult := func(r time.Time) {
				for {
					select {
					case <-done:
						return
					case <-pulse:  // 就像done channel 一样，当你执行发送或接收时，你也需要一个处理并
					// 发送心跳的 case 语句
						sendPulse()
					case results <- r:
						return
					}
				}
			}

			for {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case r := <-workGen:
					sendResult(r)
				}
			}
		}()
		return heartbeat, results
	}
```
请注意，因为我们可能在等待输入时发出多个心跳，或者在等待发送结果时发出多个心跳，所以所有的select 语句都需要在
for 循环中。目前为止看起来都很好，我们如何利用这个函数并消费它所发出的事件呢？让我们来看看：
```go
done := make(chan interface{})
	// 声明一个标准的 done channel 并在10s 后关闭。这给goroutine 留出了一些执行逻辑的时间。
	time.AfterFunc(10 * time.Second, func() {
		close(done)
	})
	// 这里我们设置了超时时间。我们使用此方法将心跳间隔与超时时间联系起来。
	const timeout = 2 * time.Second
	// 我们这里将心跳时间设为 timeout/2。这使我们的心跳有额外的响应时间，以便我们的超时不太敏感。
	heartbeat, results := doWork(done, timeout / 2)
	for {
		select {
		// 这里我们处理心跳。当没有消息时，我们至少知道每过 timeout / 2 的时间会从心跳 channel 发出一条消息。
		// 如果我们什么都没有收到，我们更知道是goroutine 本身出了问题。
		case _, ok := <-heartbeat:
			if ok == false {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results:
			if ok == false {
				return
			}
			fmt.Printf("results %v\n", r.Second())
		// 如果我们没有收到心跳或其他消息，就会超时。
		case <-time.After(timeout):
			return
		}
	}
```
[【demo】](heartbeage_goroutine/hb_goroutine.go)

输出为:
```shell
pulse
pulse
results 44
pulse
pulse
results 46
pulse
pulse
results 48
pulse
pulse
results 50
pulse
pulse
```
你可以看到，我们收到的每个消息之间大约有两个心跳。

在一个功能正常的系统中，心跳并没有什么特殊的地方。我们可能会用它们来收集关于空闲时间的统计数据，
但是当你的goroutine 不像预期的那样运行时，基于间隔的心跳的作用就会非常大。

思考下面的例子。我们将在两次迭代后停止goroutine，但却不关闭我们的任何一个channel，
来模拟一个产生了异常的goroutine。让我们看下代码：
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(
		done <-chan interface{},
		pulseInterval time.Duration,
	) (<-chan interface{}, <-chan time.Time) {
		// 我们建立了一个发送心跳的 channel 我们把这个返回给 doWork
		heartbeat := make(chan interface{})
		results := make(chan time.Time)
		go func() {

			pulse := time.Tick(pulseInterval)
			workGen := time.Tick(2 * pulseInterval)
			sendPulse := func() {
				select {
				case heartbeat <-struct{}{}:
				default:
				}
			}

			sendResult := func(r time.Time) {
				for {
					select {
					case <-pulse:
						sendPulse()
					case results <- r:
						return
					}
				}
			}
			// 这是我们模拟的问题。所以它不是无限循环的，不需要我们动手停止
			// 就像墙面的例子一样，我们只会循环两次
			for i := 0; i < 2; i ++ {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case r := <-workGen:
					sendResult(r)
				}
			}
		}()
		return heartbeat, results
	}

	done := make(chan interface{})
	time.AfterFunc(10 * time.Second, func() {
		close(done)
	})
	const timeout = 2 * time.Second
	heartbeat, results := doWork(done, timeout / 2)
	for {
		select {
		// 这里我们处理心跳。当没有消息时，我们至少知道每过 timeout / 2 的时间会从心跳 channel 发出一条消息。
		// 如果我们什么都没有收到，我们更知道是goroutine 本身出了问题。
		case _, ok := <-heartbeat:
			if ok == false {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results:
			if ok == false {
				return
			}
			fmt.Printf("results %v\n", r.Second())
		// 如果我们没有收到心跳或其他消息，就会超时。
		case <-time.After(timeout):
			fmt.Println("worker goroutine is not healthy!")
			return
		}
	}
}
```
[【demo】](hb_iter/he_iters.go)

运行得到如下:
```shell
pulse
pulse
worker goroutine is not healthy!
```
非常好，在两秒之内，我们的系统意识到我们的 goroutine 有一些不妥之处，中断了for-select 循环。
通过使用心跳，我们已经成功地避免了死锁，并且我们不需要依赖更长的超时时间来保持确定性。
我们将在本章后面 "治愈异常的goroutine" 中进一步讨论如何进一步采用这个概念。

另外请注意，心跳也会有反作用：虽然它让我们知道，长时间运行的 goroutine 依然正常工作着，但是这需要一点时间运行，
计算出值并发送给channel。

现在让我们暂时放下间隔心跳，来看看在一个工作单元开始时发出的心跳。
他对于测试来说非常有效。以下是在每个工作单元开始之前发送的例子：
```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	doWork := func(done <-chan interface{}) (<-chan interface{}, <-chan int) {
		// 创建一个缓冲区大小为1 的heartbeat channel。这确保了即使没有及时接收发送的消息，
		// 至少也会发出一个心跳。
		heartbeatStream := make(chan interface{}, 1)
		workStream := make(chan int)
		go func() {
			defer close(heartbeatStream)
			defer close(workStream)

			for i:=0; i< 10; i++ {
				// 我们为心跳设置了一个单独的select 块。我们希望将发送results 和心跳分开，
				// 因为如果接收者没有准备好接收结果，作为替代它将接收到一个心跳，而代表当前结果的值将会丢失。
				// 由于我们有默认逻辑，所以这里也没有包含对 done channel 的处理。
				select {
				case heartbeatStream <- struct{}{}:
				// 为了防止没人接收我们的心跳，我们增加了默认逻辑。因为我们的 heartbeat channel
				// 创建时有一个容量的缓冲区，所以如果有人正在监听，但是没有及时收到第一个心跳，
				// 接收者仍然可以收到心跳。
				default:
				}
				select {
				case <-done:
					return
				case workStream <- rand.Intn(10):
				}
			}
		}()
		return heartbeatStream, workStream
	}

	done := make(chan interface{})
	defer close(done)

	heartbeat, results := doWork(done)
	for {
		select {
		case _, ok := <-heartbeat:
			if ok {
				fmt.Println("pulse")
			} else {
				return
			}
		case r, ok := <-results:
			if ok {
				fmt.Printf("results %v\n", r)
			} else {
				return
			}
		}
	}
}
```
[【demo】](heart_before_work/heart_before_work.go)

运行输出:
```shell
pulse
results 1
pulse
results 7
pulse
results 7
pulse
results 9
pulse
results 1
pulse
results 8
pulse
results 5
pulse
results 0
pulse
results 6
pulse
results 0
```
你可以看到，每个结果都如预期一样伴随着一个心跳。

这种写法真正的亮点在于测试。基于时间间隔的心跳可以以相同的方式使用，
但是如果你只关心 goroutine 是否开始了它的工作，这有一种很简单的方式。
观察下面的代码:
```go
func DoWork(
	done <-chan interface{},
	nums ...int,
) (<-chan interface{}, <-chan int) {
	heartbeat := make(chan interface{}, 1)
	intStream := make(chan int)
	go func() {
		defer close(heartbeat)
		defer close(intStream)

        // 我们模拟在 goroutine 开始工作之前的某种延迟。在实践中，这可能是
        // 各种各样的问题，而且无法确定。我曾经见过CPU 负载过高、磁盘抢占、
        // 网络延迟和 goblins 造成的延迟
		time.Sleep(2 * time.Second)

		for _, n := range nums {
			select {
			case heartbeat <- struct{}{}:
			default:
			}
			select {
			case <- done:
				return
			case intStream <- n:
			}
		}
	}()

	return heartbeat, intStream
}
```
Dowork 函数是一个非常简单的生成器，它将我们传入的数字转发到它返回的channel 中。
让我们为这个函数写个测试。下面是一个不那么好的测试样例:
```go
func TestDoWorkGeneratesAllNumbers(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5}
	_, results := DoWork(done, intSlice...)

	for i, expected := range intSlice {
		select {
		case r := <-results:
			if r != expected {
				t.Errorf(
					"index %v: expected %v, but received %v,",
					i,
					expected,
					r,
				)
			}
        // 在这里我们设置一个合理的超时时间，避免 goroutine 陷入死锁
        case <-time.After(1 * time.Second):
			t.Fatal("test time out")
		}
	}
}
```
运行得到：
```go
--- FAIL: TestDoWorkGeneratesAllNumbers (1.00s)
    heartbeat_test.go:33: test time out
FAIL
exit status 1
FAIL    fzkprac/related_book_learn/Concurrency_in_go/chapter5/heartbeat/heartbeat_with_test     1.489s
```
这个测试不够好，因为它是非确定性的。虽然在我们的示例函数中，已经确保这个测试总会失败，
但如果我要删除这个time.Sleep 的话，情况可能变得更糟：这个测试有时通过，有时失败。

我们前面提到过，一些外部的因素会导致 goroutine 花费更长的时间来进行第一次迭代。无论 goroutine 在
调度中是否是第一位执行的，这都是一个令人担忧的问题。关键是我们无法保证 goroutine 的第一个迭代是否会在
超时时间结束之前执行，我们思考一下这个概率：这个超时有多大可能是有意义的？
我们可以增加超时时间，但这意味着将需要很长时间才知道执行失败，从而减慢我们的测试过程。

这会产生一些非常可怕的后果。我们慢慢开始不相信测试，然后开始忽略测试，之前的努力将一点点被瓦解。

幸运的是，利用心跳可以很轻易的解决这个问题。这是一个确定性的测试：
```go

```


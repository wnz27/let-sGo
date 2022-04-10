# reference : [技巧分享：多 Goroutine 如何优雅处理错误？](https://mp.weixin.qq.com/s?__biz=Mzg3NTU3OTgxOA==&mid=2247492387&idx=1&sn=af1b3789dad40173bbf59fd66d3d8b98&chksm=cf3df3e6f84a7af06406357cccecde5fb3915bdcf4e79bd9271d53a80932e71c05f66c9e19bd&mpshare=1&scene=1&srcid=0721k3suUvo3PzImiITh4czR&sharer_sharetime=1626856456160&sharer_shareid=d94ad27d4946e2a1fa2bda2006d8985f&version=3.1.10.90255&platform=mac#rd)

## 通过错误日志记录

为此，业务代码中常见的第一种方法：通过把错误记录写入日志文件中，再结合相关的 logtail 进行采集和梳理。 但这又会引入新的问题，那就是调用错误日志的方法写的到处都是。代码结构也比较乱，不直观。 最重要的是无法针对 error
做特定的逻辑处理和流转。

## 利用 channel 传输

这时候大家可能会想到 Go 的经典哲学：不要通过共享内存来通信，而是通过通信来实现内存共享（Do not communicate by sharing memory; instead, share memory by
communicating）。 第二种的方法：利用 channel 来传输多个 goroutine 中的 errors：

```go
package main

func main() {
	gerrors := make(chan error)
	wgDone := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		wg.Done()
	}()
	go func() {
		err := returnError()
		if err != nil {
			gerrors <- err
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(wgDone)
	}()

	select {
	case <-wgDone:
		break
	case err := <-gerrors:
		close(gerrors)
		fmt.Println(err)
	}

	time.Sleep(time.Second)
}

func returnError() error {
	return errors.New("煎鱼报错了...")
}
```

## 借助 sync/errgroup

因此第三种方法，就是使用官方提供的 sync/errgroup 标准库：

```go
type Group
func WithContext(ctx context.Context) (*Group, context.Context)
func (g *Group) Go(f func () error)
func (g *Group) Wait() error
```

- Go：启动一个协程，在新的 goroutine 中调用给定的函数。
- Wait：等待协程结束，直到来自 Go 方法的所有函数调用都返回，然后返回其中的第一个非零错误（如果有的话）。

结合其特性能够非常便捷的针对多 goroutine 进行错误处理：

```go
package main

func main() {
	g := new(errgroup.Group)
	var urls = []string{
		"http://www.golang.org/",
		"https://golang2.eddycjy.com/",
		"https://eddycjy.com/",
	}
	for _, url := range urls {
		url := url
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Printf("Errors: %+v", err)
	}
}
```

在上述代码中，其表现的是爬虫的案例。每一个计划新起的 goroutine 都直接使用 Group.Go 方法。在等待和错误上，直接调用 Group.Wait 方法就可以了。 使用标准库 sync/errgroup
这种方法的好处就是不需要关注非业务逻辑的控制代码，比较省心省力。

## 进阶使用

在真实的工程代码中，我们还可以基于 sync/errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理。以此保证能够实现一个 http server 退出，全部注销退出。
参考代码（@via 毛老师）如下：

```go
package main

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	svr := http.NewServer()
	// http server
	g.Go(
		func() error {
			fmt.Println("http")
			go func() {
				<-ctx.Done()
				fmt.Println("http ctx done")
				svr.Shutdown(context.TODO())
			}()
			return svr.Start()
		})

	// signal
	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT} // SIGTERM is POSIX specific
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		for {
			fmt.Println("signal")
			select {
			case <-ctx.Done():
				fmt.Println("signal ctx done")
				return ctx.Err()
			case <-sig:
				// do something
				return nil
			}
		}
	})

	// inject error
	g.Go(func() error {
		fmt.Println("inject")
		time.Sleep(time.Second)
		fmt.Println("inject finish")
		return errors.New("inject error")
	})

	err := g.Wait() // first error return
	fmt.Println(err)
}
```


# Goroutine 数量限制
###  M 的限制
第一，要知道在协程的执行中，真正干活的是 GPM 中的哪一个？
那势必是 M（系统线程） 了，因为 G 是用户态上的东西，最终执行都是得映射，对应到 M 这一个系统线程上去运行。
那么 M 有没有限制呢？
答案是：有的。在 Go 语言中，M 的默认数量限制是 10000，如果超出则会报错：
```shell
GO: runtime: program exceeds 10000-thread limit
```
通常只有在 Goroutine 出现阻塞操作的情况下，才会遇到这种情况。这可能也预示着你的程序有问题。
若确切是需要那么多，还可以通过 debug.SetMaxThreads 方法进行设置。

### G 的限制
第二，那 G 呢，Goroutine 的创建数量是否有限制？
答案是：没有。但理论上会受内存的影响，假设一个 Goroutine 创建需要 4k（via @GoWKH）：
4k * 80,000 = 320,000k ≈ 0.3G内存
4k * 1,000,000 = 4,000,000k ≈ 4G内存
以此就可以相对计算出来一台单机在通俗情况下，所能够创建 Goroutine 的大概数量级别。
注：Goroutine 创建所需申请的 2-4k 是需要连续的内存块。

### P 的限制
第三，那 P 呢，P 的数量是否有限制，受什么影响？
答案是：有限制。P 的数量受环境变量 GOMAXPROCS 的直接影响。
环境变量 GOMAXPROCS 又是什么？在 Go 语言中，通过设置 GOMAXPROCS，用户可以调整调度中 P（Processor）的数量。
另一个重点在于，与 P 相关联的的 M（系统线程），是需要绑定 P 才能进行具体的任务执行的，因此 P 的多少会影响到 Go 程序的运行表现。
P 的数量基本是受本机的核数影响，没必要太过度纠结他。
那 P 的数量是否会影响 Goroutine 的数量创建呢？
答案是：不影响。且 Goroutine 多了少了，P 也该干嘛干嘛，不会带来灾难性问题。

## 总结：何为之合理
在介绍完 GMP 各自的限制后，我们回到一个重点，就是 “Goroutine 数量怎么预算，才叫合理？”。
“合理” 这个词，是需要看具体场景来定义的，可结合上述对 GPM 的学习和了解。得出：
- M：有限制，默认数量限制是 10000，可调整。
- G：没限制，但受内存影响。
- P：受本机的核数影响，可大可小，不影响 G 的数量创建。
Goroutine 数量在 MG 的可控限额以下，多个把个、几十个，少几个其实没有什么影响，就可以称其为 “合理”。

## 真实情况
在真实的应用场景中，没法如此简单的定义。如果你 Goroutine：
- 在频繁请求 HTTP，MySQL，打开文件等，那假设短时间内有几十万个协程在跑，那肯定就不大合理了（可能会导致  too many files open）。
- 常见的 Goroutine 泄露所导致的 CPU、Memory 上涨等，还是得看你的 Goroutine 里具体在跑什么东西。

> 还是得看 Goroutine 里面跑的是什么东西。


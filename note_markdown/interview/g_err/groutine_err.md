reference : [技巧分享：多 Goroutine 如何优雅处理错误？](https://mp.weixin.qq.com/s?__biz=Mzg3NTU3OTgxOA==&mid=2247492387&idx=1&sn=af1b3789dad40173bbf59fd66d3d8b98&chksm=cf3df3e6f84a7af06406357cccecde5fb3915bdcf4e79bd9271d53a80932e71c05f66c9e19bd&mpshare=1&scene=1&srcid=0721k3suUvo3PzImiITh4czR&sharer_sharetime=1626856456160&sharer_shareid=d94ad27d4946e2a1fa2bda2006d8985f&version=3.1.10.90255&platform=mac#rd)

## 通过错误日志记录

为此，业务代码中常见的第一种方法：通过把错误记录写入日志文件中，再结合相关的 logtail 进行采集和梳理。 但这又会引入新的问题，那就是调用错误日志的方法写的到处都是。代码结构也比较乱，不直观。 最重要的是无法针对 error
做特定的逻辑处理和流转。

## 利用 channel 传输

这时候大家可能会想到 Go 的经典哲学：不要通过共享内存来通信，而是通过通信来实现内存共享（Do not communicate by sharing memory; instead, share memory by
communicating）。 第二种的方法：利用 channel 来传输多个 goroutine 中的 errors：

```go
func main() {
gerrors := make(chan error)
wgDone := make(chan bool)

var wg sync.WaitGroup
wg.Add(2)

go func () {
wg.Done()
}()
go func () {
err := returnError()
if err != nil {
gerrors <- err
}
wg.Done()
}()

go func () {
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
func main() {
g := new(errgroup.Group)
var urls = []string{
"http://www.golang.org/",
"https://golang2.eddycjy.com/",
"https://eddycjy.com/",
}
for _, url := range urls {
url := url
g.Go(func () error {
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






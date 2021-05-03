## 相关文章
- [X] [Goroutine Leaks - The Forgotten Sender](https://www.ardanlabs.com/blog/2018/11/goroutine-leaks-the-forgotten-sender.html)
  > goroutine 内存泄露， 可以利用缓冲channel解决。"**Never start a goroutine without knowing how it will stop**"
  - [ ] [相关阅读-context](https://blog.golang.org/context)
  - [ ] [the-behavior-of-channels](https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html)
- [X] [Concurrency Trap #2: Incomplete Work](https://www.ardanlabs.com/blog/2019/04/concurrency-trap-2-incomplete-work.html)
  > 协程泄露相关问题，使用tracker and context来配合解决
- [X] [concurrency-goroutines-and-gomaxprocs](https://www.ardanlabs.com/blog/2014/01/concurrency-goroutines-and-gomaxprocs.html)
  > 逻辑多核，并发并行不同。协程切换。不是增加核的参数就好，还是应该根据性能评价来决定 
  - [ ] [Profiling Go Programs](https://blog.golang.org/pprof)
  - [ ] [detecting-race-conditions-with-go](https://www.ardanlabs.com/blog/2013/09/detecting-race-conditions-with-go.html)
  - [ ] [video - Google IO concurrency youtube](https://www.youtube.com/watch?v=f6kdp27TYZs)
- [X] [concurrency](https://dave.cheney.net/practical-go/presentations/qcon-china.html#_concurrency)
  - Keep yourself busy or do the work yourself, 不要过度使用goroutine
  - Leave concurrency to the caller, 解耦 调用者和 异步函数（goroutine）来控制goroutine的执行
  - Never start a goroutine without knowing how it will stop， 防止泄露，利用缓冲channel
- [X] [The Go Memory Model](https://golang.org/ref/mem)
  - A receive from an unbuffered channel happens before the send on that channel completes.
  - The kth receive on a channel with capacity C happens before the k+Cth send from that channel completes.
  - For any sync.Mutex or sync.RWMutex variable l and n < m, call n of l.Unlock() happens before call m of l.Lock() returns.
  - For any call to l.RLock on a sync.RWMutex variable l, there is an n such that the l.RLock happens (returns) after call n to l.Unlock and the matching l.RUnlock happens before call n+1 to l.Lock.
  - A single call of f() from once.Do(f) happens (returns) before any call of once.Do(f) returns.
  - go 内存模型
	- 1、w happens before r.
	- 2、Any other write to the shared variable v either happens before w or after r.
- [X] [理解Memory Barrier（内存屏障）](https://blog.csdn.net/caoshangpa/article/details/78853919)
  - [相关文章](https://blog.csdn.net/world_hello_100/article/details/50131497)
  - 编译时内存乱序访问
  ```
  // thread 1
  while (!ok);
  do(x);
  
  // thread 2
  x = 42;
  ok = 1;
  ```
  > 线程2两条写语句顺序不定，可能导致do的时候x不是希望的42
  - 运行时内存乱序访问
  - Memory Barrier

- [X] [内存重排](https://blog.csdn.net/qcrao/article/details/92759907)
  > 再高速缓存还没有到内存，可能导致重排。
  A barrier instruction forces all memory operations before it to complete before any memory operation after it can begin.
  barrier 指令要求所有对内存的操作都必须要“扩散”到 memory 之后才能继续执行其他对 memory 的操作。
  正是 CPU 提供的 barrier 指令，我们才能实现应用层的各种同步原语，如 atomic，而 atomic 又是各种更上层的 lock 的基础。
  - [ ] [memory_barrier](https://github.com/cch123/golang-notes/blob/master/memory_barrier.md)
- [ ] [从 Memory Reordering 说起](https://cch123.github.io/ooo/)
https://blog.golang.org/codelab-share
https://dave.cheney.net/2018/01/06/if-aligned-memory-writes-are-atomic-why-do-we-need-the-sync-atomic-package
http://blog.golang.org/race-detector
https://dave.cheney.net/2014/06/27/ice-cream-makers-and-data-races
https://www.ardanlabs.com/blog/2014/06/ice-cream-makers-and-data-races-part-ii.html
https://medium.com/a-journey-with-go/go-how-to-reduce-lock-contention-with-the-atomic-package-ba3b2664b549
https://medium.com/a-journey-with-go/go-discovery-of-the-trace-package-e5a821743c3c
https://medium.com/a-journey-with-go/go-mutex-and-starvation-3f4f4e75ad50
https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html
https://medium.com/a-journey-with-go/go-buffered-and-unbuffered-channels-29a107c00268
https://medium.com/a-journey-with-go/go-ordering-in-select-statements-fd0ff80fd8d6
https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html
https://www.ardanlabs.com/blog/2014/02/the-nature-of-channels-in-go.html
https://www.ardanlabs.com/blog/2013/10/my-channel-select-bug.html
https://blog.golang.org/io2013-talk-concurrency
https://blog.golang.org/waza-talk
https://blog.golang.org/io2012-videos
https://blog.golang.org/concurrency-timeouts
https://blog.golang.org/pipelines
https://www.ardanlabs.com/blog/2014/02/running-queries-concurrently-against.html
https://blogtitle.github.io/go-advanced-concurrency-patterns-part-3-channels/
https://www.ardanlabs.com/blog/2013/05/thread-pooling-in-go-programming.html
https://www.ardanlabs.com/blog/2013/09/pool-go-routines-to-process-task.html
https://blogtitle.github.io/categories/concurrency/
https://medium.com/a-journey-with-go/go-context-and-cancellation-by-propagation-7a808bbc889c
https://blog.golang.org/context
https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html
https://golang.org/ref/spec#Channel_types
https://drive.google.com/file/d/1nPdvhB0PutEJzdCq5ms6UI58dp50fcAN/view
https://medium.com/a-journey-with-go/go-context-and-cancellation-by-propagation-7a808bbc889c
https://blog.golang.org/context
https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html
https://golang.org/doc/effective_go.html#concurrency
https://zhuanlan.zhihu.com/p/34417106?hmsr=toutiao.io
https://talks.golang.org/2014/gotham-context.slide#1
https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39


## 书籍
- [ ] 《Go语言并发之道》Katherine 著，中国电力出版社

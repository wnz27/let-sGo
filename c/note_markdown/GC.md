# go GC

gColl.go的􏱑一􏰐分代􏰅段􏰽下所􏰝:
```go
package main
import (
    "fmt"
    "os"
    "runtime"
    "runtime/trace"
    "time"
)
func printStats(mem runtime.MemStats) {
    runtime.ReadMemStats(&mem)
    fmt.Println("mem.Alloc:", mem.Alloc)
    fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
    fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
    fmt.Println("mem.NumGC:", mem.NumGC)
    fmt.Println("-----")
}
```

这里有一个技􏰪巧可以让你得到更多关于go 垃圾收集器操作的细节，使用 下面这个命令。
```
GODEBUG=gctrace=1 go run gColl.go
```
所以，如果你在任何 `go run` 命令前面加上 `GODEBUG=gctrace=1` ，go就会去打印关于垃圾回收操作的一些分析数据。
长这样：
```go
gc 4 @0.025s 0%: 0.002+0.65+0.018 ms clock, 0.021+0.040/0.057/0.
gc 17 @30.103s 0%: 0.004+0.080+0.019ms clock, 0.033+0/0.076/0.07
```
这些数据给你提供了更多垃圾回收过程中的堆内存大的信息。
让我们 以 47->47->0 MB 这三个值为例。
第一个垃圾回收器要去运行时候的堆内存大小。
第二个值是垃圾回收器操作结束时候的堆内存大小。
最后一个值就是生存堆的大小。





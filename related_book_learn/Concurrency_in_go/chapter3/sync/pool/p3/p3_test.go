/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/9 23:50 8月
 **/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

/*
另一种常见的情况是，用Pool来尽可能快地将预先分配的对象缓存加载启动。
在这种情况下，我们不是试图通过限制创建的对象的数量来节省主机的内存，
而是通过提前加载获取引用到另一个对象所需的时间，来节省消费者的时间。

这在编写高吞吐量网络服务器时十分常见，服务器视图快速响应请求。
让我们看看这样的场景。
 */

// 首先让我们创建一个模拟创建到服务的连接的函数，我们会让这个连接话费很长时间。

func connectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct {}{}
}

/*
	接下来，让我们了解一下， 如果服务为每个请求都启动一个新的连接，那么网络服务的性能如何。
	我们将编写一个网络处理程序，为每个请求都打开一个新的连接。
	为了使基准测试检点，我们只允许一次连接：
*/
func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}
			connectToService()
			fmt.Fprintln(conn, "")
			conn.Close()
		}
	}()
	return &wg
}

// 现在我们的基准如下：
func init() {
	daemonStarted := startNetworkDaemon()
	daemonStarted.Wait()
}

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}

// cd 到这个文件所在文件夹下
// 执行 go test -benchtime=10s -bench=.

/*
我的机器输出：
goos: darwin
goarch: amd64
pkg: fzkprac/related_book_learn/Concurrency_in_go/chapter3/sync/pool/p3
cpu: Intel(R) Core(TM) i7-4870HQ CPU @ 2.50GHz
BenchmarkNetworkRequest-8             10        1004891752 ns/op
PASS
ok      fzkprac/related_book_learn/Concurrency_in_go/chapter3/sync/pool/p3      11.442s

也就是1秒出头一次操作，这在我们设置的来看是合理的，但是我们看看是否可以通过使用sync.Pool来改进我们的虚拟服务

示例见下一个demo
 */




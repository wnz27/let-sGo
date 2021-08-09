/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/10 00:54 8月
 **/
package pool

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
)

/*
尝试用池化技术来改进我们虚拟服务
 */

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}
	for i := 0; i < 10; i ++ {
		p.Put(p.New())
	}
	return p
}

func startNetworkDaemon2() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		connPool := warmServiceConnCache()
		server, err := net.Listen("tcp", "localhost:8181")
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
			svcConn := connPool.Get()
			fmt.Fprintln(conn, "")
			connPool.Put(svcConn)
			conn.Close()
		}
	}()
	return &wg
}

// 现在我们的基准如下：
func init() {
	daemonStarted := startNetworkDaemon2()
	daemonStarted.Wait()
}

func BenchmarkNetworkRequest2(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		conn, err := net.Dial("tcp", "localhost:8181")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}


// 继续使用  go test -benchtime=10s -bench=.
/*
结果：
goos: darwin
goarch: amd64
pkg: fzkprac/related_book_learn/Concurrency_in_go/chapter3/sync/pool/p3
cpu: Intel(R) Core(TM) i7-4870HQ CPU @ 2.50GHz
BenchmarkNetworkRequest-8             10        1005063395 ns/op
BenchmarkNetworkRequest2-8          2437           7997014 ns/op
PASS
ok      fzkprac/related_book_learn/Concurrency_in_go/chapter3/sync/pool/p3      47.884s

可以从数值看到快了三个数量级！你可以看到在处理待加昂贵的事务时使用这种模式可以极大地提高响应时间。
 */


/*
你的并发进程需要请求一个对象，但是在实例化之后很快的处理它们时，或者在这些对象的构造可能会对内存产生负面影响，
这时最好使用Pool设计模式
 */

// Todo 总结
/*
有些情况下要谨慎决定你是否应该使用Pool：如果你使用Pool代码所需要的东西不是大概同质的，
那么从Pool中转化检索所需要的内容的时间可能比重新实例化内容要话费的时间更多。

例如，如果你的程序需要随机和可变长度的切片，那么Pool将不会对你有多大的帮助。
你直接从Pool中获得一个正确的切片的概率是很低的。
 */

// TODO 注意事项
/*
1、 当实例化sync.Pool, 使用new方法创建一个成员变量，在调用时是线程安全的。
2、 当你收到一个来自Get的实例时，不要对所接受的对象的状态做出任何假设。todo: 疑问
3、 当你用完了一个从Pool中取出来的对象时，一定要调用Put，否则，Pool就无法复用这个实例了。
通常情况下，这是用defer完成的
4、 Pool 内的分布必须大致均匀
 */

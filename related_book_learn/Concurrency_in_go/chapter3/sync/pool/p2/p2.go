/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/9 23:26 8月
 **/
package main

import (
	"fmt"
	"sync"
)

/*
为什么要使用Pool， 而不知识在运行时实例化对象呢？
Go 语言是有GC的，因此实例化的对象将被自动清除。有什么意义？
思考下面的例子
 */

func main() {
	var numCalcsCreated int
	calcPool := &sync.Pool {
		New: func () interface{} {
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	// 用4KB 初始化pool
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i -- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)

			// 做一些有趣的假设，但是很快就会用这个内存完成
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)

	/*
	输出
	8 calculators were created.

	如果我没有用sync.Pool运行这个例子，尽管结果是不确定的， 在最坏的情况下，
	可能尝试分配一个十亿字节的内存，但是正如从输出看到的，我只分配了4KB。

	 */

}

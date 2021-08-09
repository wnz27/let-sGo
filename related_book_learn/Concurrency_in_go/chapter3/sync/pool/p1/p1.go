/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/9 02:10 8月
 **/
package main

import (
	"fmt"
	"sync"
)

/*
池是Pool模式的并发安全实现。
在较高的层次上，Pool模式是一种创建和提供可供使用的固定数量的实例或Pool实例的方法。
它通常用于约束创建昂贵的场景（如数据库连接），以便只创建固定数量的实例，
但不确定数量的操作仍然可以请求访问这些场景。

！！！！对于Go语言的sync.Pool，这种数据类型可以被多个goroutine安全地使用。

Pool的主接口是它的Get方法。当调用时，Get将首先检查池中是否有可用的实例返回给调用者，
如果没有，调用它的new方法来创建一个新实例。
当完成时，调用者调用Put方法把工作的实例归还到池中，以供其他进程使用。
 */
func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}
	myPool.Get()
	instance := myPool.Get()
	// 调用Pool的get方法，这些调用将执行Pool中的定义的new函数，因为实例还没有实例化。

	// 我们将先前检索到的实例放在池中, 增加了实例的可用数量
	myPool.Put(instance)
	//todo 自己思考： 这里意味着要显式的放回?

	fmt.Println("===================================")
	// 在执行此调用时，我们将重用以前分配的实例并将其放回池中。New将不会被调用。
	myPool.Get()

	// 输出:
	/*
	Creating new instance.
	Creating new instance.
	===================================
	 */
}

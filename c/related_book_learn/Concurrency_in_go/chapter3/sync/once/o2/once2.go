/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/9 01:50 8月
 **/
package main

import (
	"fmt"
	"sync"
)

func main() {
	// 来看另一个例子
	var count int
	increment := func() {count ++}
	decrement := func() {count --}
	var once sync.Once
	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("Count %d \n", count)
	/*
	这就可能会惊讶了， 输出的是1 而不是0， 这是因为
	sync.Once只计算调用Do的次数，而不是多少次Do唯一调用的方法。
	这样，sync.Once 的副本与索要调用的函数紧密耦合，
	我们再次看到如何在一个严格的范围内合理的使用sync包中的类型以发挥最佳效果。

	建议通过将sync.Once 包装在一个小的语法块中来形式化这种耦合：
	要么是一个小函数，要么是将两者包装在一个结构体中。
	 */

}

/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/2 3:16 上午
* Description:
 */
package main

import (
	"fmt"
	"sync/atomic"
)

// 竞态检测——检测代码在并发环境下可能出现的问题
// 当多线程并发运行的程序竞争访问和修改同一块资源时，会发生竞态问题。
// 下面的代码中有一个ID生成器，每次调用生成器将会生成一个不会重复的顺序序号，使用10个并发生成序号，观察10个并发后的结果。
var (
	// 序列号
	seq int64
)

// 序列号生成器, 有误
func GenID() int64 {
	// 尝试原子的增加序列号
	atomic.AddInt64(&seq, 1)
	return seq
}

func GenIDRight() int64 {
	// 尝试原子的增加序列号
	return atomic.AddInt64(&seq, 1)
}


// go run -race xxx.go 来检测竞态条件
func main(){
	// 生成10个并发序列号
	for i := 0; i < 10 ; i ++ {
		go GenIDRight()
	}

	fmt.Println(GenIDRight())
}

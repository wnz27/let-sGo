/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/5/2 09:44 5月
 **/
package main

import (
	"fmt"
	"sync"
)

// 互斥锁是一种常用的控制共享资源访问的方法, 在go程序中使用非常简单，参见下面的代码：
var (
	// 逻辑中使用的某个变量
	count int

	// 与变量对应的使用互斥锁
	countGuard sync.Mutex
)

func GetCount() int {
	// 锁定
	countGuard.Lock()

	// 在函数退出时解除锁定
	defer countGuard.Unlock()

	return count
}

func setCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}


func main(){
	// 可以进行并发安全性设置
	setCount(1)
	// 可以进行并发安全的获取
	fmt.Println(GetCount())


}

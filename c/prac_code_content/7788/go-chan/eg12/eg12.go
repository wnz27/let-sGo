/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/5/2 09:52 5月
 **/
package main

import (
	"fmt"
	"sync"
)

var (
	// 逻辑中使用的某个变量
	count int

	// 与变量对应的使用互斥锁
	countGuard sync.RWMutex

)

func GetCount() int {
	// 锁定
	countGuard.RLock()

	// 在函数退出时解锁
	defer countGuard.RUnlock()

	return count
}

func setCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}


func main(){
	// 可以进行并发安全性设置
	setCount(29292)
	// 可以进行并发安全的获取
	fmt.Println(GetCount())

}


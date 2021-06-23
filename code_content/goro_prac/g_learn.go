/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/6/23 11:37 下午
* Description:
 */

/*
实现控制并发的方式，大致可分成以下三类：

全局共享变量

channel通信

Context包
 */

/*
todo 1、这是最简单的实现控制并发的方式，实现步骤是：
	- 声明一个全局变量；
	- 所有子goroutine共享这个变量，并不断轮询这个变量检查是否有更新；
	- 在主进程中变更该全局变量；
	- 子goroutine检测到全局变量更新，执行相应的逻辑。
*/
package main

func main() {

}



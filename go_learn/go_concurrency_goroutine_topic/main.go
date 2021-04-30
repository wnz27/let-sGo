/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/4/30 23:28 4月
 **/
package main

import (
	"fmt"
	"time"
)

func slowQueryOrOtherConsumeTimeOperation0() {
	time.Sleep(time.Second * 2)
	fmt.Println("slowQueryOrOtherConsumeTimeOperation finished!")
}

func slowQueryOrOtherConsumeTimeOperation(c chan<- string) {
	time.Sleep(time.Second * 2)
	c <- "slowQueryOrOtherConsumeTimeOperation finished!"
}

func main(){
	// v0
	go slowQueryOrOtherConsumeTimeOperation0()
	fmt.Println("slowQueryOrOtherConsumeTimeOperation0  not execute!")
	fmt.Println("// ---------------------- ---  ---------------------- //")
	// v1
	slowDatas := make(chan string)
	go slowQueryOrOtherConsumeTimeOperation(slowDatas)

	// 在慢速的玩意儿整完之前我可以先做其他的
	fmt.Println("do others")

	// 等待通道信息
	msg :=  <- slowDatas
	fmt.Println(msg)

	fmt.Println("// ---------------------- ---  ---------------------- //")

	// v2

}

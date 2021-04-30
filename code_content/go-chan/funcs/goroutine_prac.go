/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/4/30 21:59 4月
 **/
package funcs

import (
	"fmt"
	"time"
)

func Tg() {
	c := make(chan string)
	c <- "hello world"
	msg := <-c
	fmt.Println(msg)
}


func slowFunc(c chan<- string){
	time.Sleep(time.Second * 2)
	c <- "slowFunc() finished!"
}

func OutFunc() string {
	c := make(chan string)
	go slowFunc(c)

	return <-c
}



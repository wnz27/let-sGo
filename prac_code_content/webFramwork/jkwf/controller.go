/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-26 11:07:29
 * @LastEditTime: 2022-01-26 16:32:10
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/controller.go
 * @description: type some description
 */

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"jkframe"
)

func FooControllerHandler(c *jkframe.Context) error {
	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
	// 这里记得当所有事情处理结束后调用 cancel，告知 durationCtx 的后续 Context 结束
	defer cancel()
	// 这个 channal 负责通知结束
	finish := make(chan struct{}, 1)
	// 这个 channel 负责通知 panic 异常
	panicChan := make(chan interface{}, 1)

	go func() {
		// 这里增加异常处理
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// 这里做具体的业务
		time.Sleep(10 * time.Second)
		c.SetOkStatus().Json("ok")
		// 新的 goroutine 结束的时候通过一个 finish 通道告知父 goroutine
		finish <- struct{}{}
	}()

	// 请求监听的时候增加锁机制
	select {
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.SetStatus(500).Json("panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.SetStatus(500).Json("time out")
		c.SetHasTimeout()
	}
	return nil
}

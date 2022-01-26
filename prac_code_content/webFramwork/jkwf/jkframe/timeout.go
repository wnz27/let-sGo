/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-26 15:14:10
 * @LastEditTime: 2022-01-26 15:14:10
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/jkframe/timeout.go
 * @description: type some description
 */

package jkframe

import (
	"context"
	"fmt"
	"log"
	"time"
)

func TimeoutHandler(fun ControllerHandler, d time.Duration) ControllerHandler {
	// 使用函数回调
	return func(c *Context) error {

		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		// 执行业务逻辑前预操作：初始化超时 context
		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

		c.request.WithContext(durationCtx)

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			// 执行具体的业务逻辑
			fun(c)

			finish <- struct{}{}
		}()
		// 执行业务逻辑后操作
		select {
		case p := <-panicChan:
			log.Println(p)
			c.responseWriter.WriteHeader(500)
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.SetHasTimeout()
			c.responseWriter.Write([]byte("time out"))
		}
		return nil
	}
}

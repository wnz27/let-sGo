/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-26 15:46:09
 * @LastEditTime: 2022-01-26 15:46:09
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/jkframe/middleware/timeout.go
 * @description: type some description
 */

package middleware

import (
	"context"
	"fmt"
	"jkframe"
	"log"
	"time"
)

func Timeout(d time.Duration) jkframe.ControllerHandler {
	// 使用函数回调
	return func(c *jkframe.Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)
		// 执行业务逻辑前预操作：初始化超时context
		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			// 使用next执行具体的业务逻辑
			c.Next()

			finish <- struct{}{}
		}()
		// 执行业务逻辑后操作
		select {
		case p := <-panicChan:
			c.SetStatus(500).Json("time out")
			log.Println(p)
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.SetHasTimeout()
			c.SetStatus(500).Json("time out")
		}
		return nil
	}
}

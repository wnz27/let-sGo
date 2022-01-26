/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-26 16:10:39
 * @LastEditTime: 2022-01-26 16:10:40
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/jkframe/middleware/cost.go
 * @description: type some description
 */

package middleware

import (
	"log"
	"time"

	"jkframe"
)

// recovery机制，将协程中的函数异常进行捕获
func Cost() jkframe.ControllerHandler {
	// 使用函数回调
	return func(c *jkframe.Context) error {
		// 记录开始时间
		start := time.Now()

		// 使用next执行具体的业务逻辑
		c.Next()

		// 记录结束时间
		end := time.Now()
		cost := end.Sub(start)
		log.Printf("api uri: %v, cost: %v", c.GetRequest().RequestURI, cost.Seconds())

		return nil
	}
}

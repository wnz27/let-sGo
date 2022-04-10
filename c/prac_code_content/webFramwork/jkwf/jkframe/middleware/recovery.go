/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-26 15:46:45
 * @LastEditTime: 2022-01-26 16:33:29
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/jkframe/middleware/recovery.go
 * @description: type some description
 */

package middleware

import (
	"jkframe"
)

// recovery机制，将协程中的函数异常进行捕获
func Recovery() jkframe.ControllerHandler {
	// 使用函数回调
	return func(c *jkframe.Context) error {
		// 核心在增加这个recover机制，捕获c.Next()出现的panic
		defer func() {
			if err := recover(); err != nil {
				c.SetStatus(500).Json(err)
			}
		}()
		// 使用next执行具体的业务逻辑
		c.Next()

		return nil
	}
}

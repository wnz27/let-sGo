/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-26 11:11:06
 * @LastEditTime: 2022-01-26 15:48:51
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/route.go
 * @description: type some description
 */

package main

import (
	"jkframe"
	"jkframe/middleware"
)

// 注册路由规则
func registerRouter(core *jkframe.Core) {
	// 在core中使用middleware.Test3() 为单个路由增加中间件
	core.Get("/user/login", middleware.Test3(), UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		// 在group中使用middleware.Test3() 为单个路由增加中间件
		subjectApi.Get("/:id", middleware.Test3(), SubjectGetController)
	}
}

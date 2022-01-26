/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-26 11:11:06
 * @LastEditTime: 2022-01-26 15:08:35
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/route.go
 * @description: type some description
 */

package main

import "jkframe"

// 注册路由规则
func registerRouter(core *jkframe.Core) {
	// 需求1+2:HTTP方法+静态路由匹配
	core.Get("/user/login", UserLoginController)

	// 需求3:批量通用前缀
	subjectApi := core.Group("/subject")
	{
		// 需求4:动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)
	}
}

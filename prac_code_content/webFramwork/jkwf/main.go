/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-21 15:07:29
 * @LastEditTime: 2022-01-26 15:44:37
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/main.go
 * @description: type some description
 */

package main

import (
	"net/http"

	"jkframe"
	"jkframe/middleware"
)

func main() {
	jkframe.PrintName()
	core := jkframe.NewCore()

	// core中使用use注册中间件
	core.Use(
		middleware.Test1(),
		middleware.Test2())

	// group中使用use注册中间件
	subjectApi := core.Group("/subject")
	subjectApi.Use(middleware.Test3())
	// 设置路由
	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: core,
		// 请求监听地址
		Addr: ":8888",
	}
	server.ListenAndServe()
}

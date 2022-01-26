/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-21 15:07:29
 * @LastEditTime: 2022-01-26 11:18:57
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/main.go
 * @description: type some description
 */

package main

import (
	"net/http"

	"jkframe"
)

func main() {
	jkframe.PrintName()
	core := jkframe.NewCore()
	// 设置路由
	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: core,
		// 请求监听地址
		Addr: ":8888",
	}
	server.ListenAndServe()
}

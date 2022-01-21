/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-21 15:07:29
 * @LastEditTime: 2022-01-21 17:47:30
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
	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: jkframe.NewCore(),
		// 请求监听地址
		Addr: ":8080",
	}
	server.ListenAndServe()
}

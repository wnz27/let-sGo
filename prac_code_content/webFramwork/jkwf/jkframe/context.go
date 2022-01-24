/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-24 17:42:21
 * @LastEditTime: 2022-01-24 17:44:16
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/jkframe/context.go
 * @description: type some description
 */

package jkframe

import "net/http"

type Context struct {
	// 获取请求、返回结果功能
	request        *http.Request
	responseWriter http.ResponseWriter
}

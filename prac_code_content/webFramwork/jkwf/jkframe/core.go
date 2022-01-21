/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-21 17:37:40
 * @LastEditTime: 2022-01-21 17:42:00
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/jkframe/core.go
 * @description: type some description
 */

package jkframe

import (
	"fmt"
	"net/http"
)

func PrintName() {
	fmt.Println("jkframe name")
}

// 框架核心结构
type Core struct {
}

// 初始化框架核心结构
func NewCore() *Core {
	return &Core{}
}

// 框架核心结构实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// TODO
}

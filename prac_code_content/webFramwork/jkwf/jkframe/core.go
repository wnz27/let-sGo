/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-21 17:37:40
 * @LastEditTime: 2022-01-26 11:19:40
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/jkframe/core.go
 * @description: type some description
 */

package jkframe

import (
	"fmt"
	"log"
	"net/http"
)

func PrintName() {
	fmt.Println("jkframe name")
}

// 框架核心结构
type Core struct {
	router map[string]ControllerHandler
}

// 初始化框架核心结构
func NewCore() *Core {
	return &Core{router: map[string]ControllerHandler{}}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

// 框架核心结构实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core.serveHTTP")
	ctx := NewContext(request, response)

	// 一个简单的路由选择器，这里直接写死为测试路由foo
	router := c.router["foo"]
	if router == nil {
		return
	}
	log.Println("core.router")

	router(ctx)
}

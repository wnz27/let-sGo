/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-26 15:06:14
 * @LastEditTime: 2022-01-26 17:40:44
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/user_controller.go
 * @description: type some description
 */

package main

import (
	"time"

	"jkframe"
)

func UserLoginController(c *jkframe.Context) error {
	foo, _ := c.QueryString("foo", "def")
	// 等待10s才结束执行
	time.Sleep(10 * time.Second)
	// 输出结果
	c.SetOkStatus().Json("ok, UserLoginController: " + foo)
	return nil
}

/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-26 15:06:14
 * @LastEditTime: 2022-01-26 15:06:15
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/user_controller.go
 * @description: type some description
 */

package main

import "jkframe"

func UserLoginController(c *jkframe.Context) error {
	// 打印控制器名字
	c.Json(200, "ok, UserLoginController")
	return nil
}

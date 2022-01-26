/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-26 11:11:06
 * @LastEditTime: 2022-01-26 11:34:06
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/route.go
 * @description: type some description
 */

package main

import "jkframe"

func registerRouter(core *jkframe.Core) {
	// core.Get("foo", framework.TimeoutHandler(FooControllerHandler, time.Second*1))
	core.Get("foo", FooControllerHandler)
}

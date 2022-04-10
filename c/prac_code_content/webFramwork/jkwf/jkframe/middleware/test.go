/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-26 15:37:40
 * @LastEditTime: 2022-01-26 15:45:34
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/jkframe/middleware/test.go
 * @description: type some description
 */

package middleware

import (
	"fmt"

	"jkframe"
)

func Test1() jkframe.ControllerHandler {
	// 使用函数回调
	return func(c *jkframe.Context) error {
		fmt.Println("middleware pre test1")
		c.Next() // 调用Next往下调用，会自增contxt.index
		fmt.Println("middleware post test1")
		return nil
	}
}

func Test2() jkframe.ControllerHandler {
	// 使用函数回调
	return func(c *jkframe.Context) error {
		fmt.Println("middleware pre test2")
		c.Next() // 调用Next往下调用，会自增contxt.index
		fmt.Println("middleware post test2")
		return nil
	}
}

func Test3() jkframe.ControllerHandler {
	// 使用函数回调
	return func(c *jkframe.Context) error {
		fmt.Println("middleware pre test3")
		c.Next()
		fmt.Println("middleware post test3")
		return nil
	}
}

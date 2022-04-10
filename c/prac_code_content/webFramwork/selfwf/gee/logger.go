/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-21 17:03:53
 * @LastEditTime: 2022-01-21 17:03:53
 * @FilePath: /let-sGo/prac_code_content/webFramwork/selfwf/gee/logger.go
 * @description: type some description
 */

package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

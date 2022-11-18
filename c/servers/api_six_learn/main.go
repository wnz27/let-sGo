/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-11-18 11:45:57
 * @LastEditTime: 2022-11-18 12:47:31
 * @FilePath: /let-sGo/c/servers/api_six_learn/main.go
 * @description: type some description
 */

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong111111",
		})
	})
	r.Run(":8888") // 监听并在 0.0.0.0:8888 上启动服务
}

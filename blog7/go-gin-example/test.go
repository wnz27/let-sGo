/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/6/16 4:01 下午
* Description:
 */
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":9090") // listen and serve on 0.0.0.0:9090
}
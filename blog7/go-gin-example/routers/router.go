/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/6/16 10:34 下午
* Description:
 */
package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	return r
}

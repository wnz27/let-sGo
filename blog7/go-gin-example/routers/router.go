/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/6/16 10:34 下午
* Description:
 */
package routers

import (
	v1 "github.com/EDDYCJY/go-gin-example/routers/api/v1"
	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
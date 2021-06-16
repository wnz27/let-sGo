/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/6/16 4:01 下午
* Description:
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run(":9090") // listen and serve on 0.0.0.0:9090
	b1 := 1 << 20 //  1 乘以 2的20次方
	b := int(math.Pow(2, 20))
	fmt.Println("b", " = ", b, "b1 = ", b1)
	a:=12
	fmt.Printf("%d\n",a<<2)
}
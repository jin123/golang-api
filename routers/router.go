package routers

import (
	"golang-api/controllers"

	"github.com/gin-gonic/gin"
)

func StartRouters() *gin.Engine {
	r := gin.Default()
	test := &controllers.TestCtl{}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/test", test.Test1())
	return r
}

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is test1111",
	})
}

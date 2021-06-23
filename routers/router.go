package routers

import (
	"fmt"
	"golang-api/controllers"

	"github.com/gin-gonic/gin"
)

//--------------------------------------------//
type MyContext struct {
	*gin.Context
	userId int64
}

//
func do(c *MyContext) {
	fmt.Println(c.userId) //123456
}

//
type HandlerFunc func(*MyContext)

//
func MyHandler(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		//
		ctx := new(MyContext)
		ctx.Context = c

		//
		handler(ctx)
	}
}

//--------------------------------------------//
//
//--------------------------------------------//
func MyMiddleware(ctx *gin.Context) {
	ctx.Set("mykey", 10)
	ctx.Set("mykey2", "m1")
	ctx.Set("keyUserId", int64(123456)) //
}

//--------------------------------------------//

func StartRouters() *gin.Engine {
	r := gin.Default()
	test := &controllers.TestCtl{}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Use(gin.Recovery())
	r.GET("/test", test.Test1)
	r.GET("/hi", MyMiddleware, MyHandler(do))
	return r
}

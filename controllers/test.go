package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type TestCtl struct {
	ctl
}

func (t *TestCtl) Test1(c *gin.Context) {
	fmt.Println(" this is ctl")
	userId, ok := c.Get("keyUserId")
	fmt.Println(userId, ok)
	c.JSON(200, gin.H{
		"message": "this is ctl test",
	})
}

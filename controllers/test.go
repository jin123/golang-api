package controllers

import "github.com/gin-gonic/gin"

type TestCtl struct {
	ctl
}

func (t *TestCtl) Test1(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is ctl test",
	})
}

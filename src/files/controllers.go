package files

const ControllersMain = `
package controllers

import (
	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	context.JSON(200, gin.H{"message": "ping"})
}
`

const ControllersUtilities = `
package controllers

import (
	"github.com/gin-gonic/gin"
)

func status500(c *gin.Context, err error) {
	c.JSON(500, gin.H{
		"message": "An internal error occurred",
		"data":    err.Error(),
	})
}
`

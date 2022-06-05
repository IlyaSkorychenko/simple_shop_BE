package api

import (
	"github.com/IlyaSkorychenko/simple_shop_BE/pkg"
	"github.com/gin-gonic/gin"
)

func errHandler(c *gin.Context) {
	c.Next()

	for _, ginErr := range c.Errors {
		switch err := ginErr.Err.(type) {
		case pkg.IHttpError:
			c.AbortWithStatusJSON(err.GetCode(), gin.H{"message": err.Error()})
		default:
			c.AbortWithStatusJSON(500, gin.H{"message": "Unknown error"})
		}
	}
}

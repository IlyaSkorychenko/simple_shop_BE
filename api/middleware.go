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
			response := gin.H{"message": err.Error()}

			if err.Errors() != nil {
				response = gin.H{
					"message": err.Error(),
					"errors":  err.Errors(),
				}
			}

			c.AbortWithStatusJSON(err.GetCode(), response)
		default:
			c.AbortWithStatusJSON(500, gin.H{"message": "Unknown error"})
		}
	}
}

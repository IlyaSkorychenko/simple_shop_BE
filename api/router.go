package api

import (
	"github.com/IlyaSkorychenko/simple_shop_BE/api/handler"
	"github.com/gin-gonic/gin"
)

func (s Server) addProductRouts(rg *gin.RouterGroup) {
	product := rg.Group("/products")

	product.GET("/", handler.GetProducts)
}

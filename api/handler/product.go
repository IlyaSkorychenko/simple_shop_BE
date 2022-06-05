package handler

import (
	"github.com/IlyaSkorychenko/simple_shop_BE/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProducts(c *gin.Context) {
	products, err := service.GetProducts()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, products)
}

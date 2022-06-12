package handler

import (
	"github.com/IlyaSkorychenko/simple_shop_BE/pkg"
	"github.com/IlyaSkorychenko/simple_shop_BE/pkg/entity"
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

func CreateProduct(c *gin.Context) {
	requestBody, exist := c.Get("body")
	if !exist {
		c.Error(pkg.CustomInternalServerError("getting request error"))
		return
	}

	dto := requestBody.(entity.ProductDto)
	err := service.CreateProduct(dto)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

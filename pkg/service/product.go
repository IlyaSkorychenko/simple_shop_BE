package service

import (
	"github.com/IlyaSkorychenko/simple_shop_BE/pkg"
	"github.com/IlyaSkorychenko/simple_shop_BE/pkg/entity"
	"github.com/IlyaSkorychenko/simple_shop_BE/pkg/repository"
)

func GetProducts() ([]entity.Product, *pkg.HttpError) {
	return repository.GetAllProducts()
}

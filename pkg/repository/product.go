package repository

import (
	"github.com/IlyaSkorychenko/simple_shop_BE/database"
	"github.com/IlyaSkorychenko/simple_shop_BE/pkg"
	ent "github.com/IlyaSkorychenko/simple_shop_BE/pkg/entity"
)

func GetAllProducts() ([]ent.Product, *pkg.HttpError) {
	rows, err := database.GoquDb.
		From("destination").
		Select("*").
		Executor().
		Query()
	if err != nil {
		return nil, pkg.NewInternalServerError(err, "DB error")
	}

	var products []ent.Product

	for rows.Next() {
		var product ent.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
			return products, pkg.NewInternalServerError(err, "Mapping struct error")
		}
		products = append(products, product)
	}

	return products, nil
}

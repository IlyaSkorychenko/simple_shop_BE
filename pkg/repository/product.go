package repository

import (
	"fmt"
	"github.com/IlyaSkorychenko/simple_shop_BE/database"
	"github.com/IlyaSkorychenko/simple_shop_BE/pkg"
	ent "github.com/IlyaSkorychenko/simple_shop_BE/pkg/entity"
	"github.com/doug-martin/goqu/v9"
)

func GetAllProducts() ([]ent.Product, *pkg.HttpError) {
	rows, err := database.GoquDb.
		From(database.ProductTable).
		Select("*").
		Executor().
		Query()
	if err != nil {
		return nil, pkg.InternalServerError(err, "DB error")
	}

	var products []ent.Product

	for rows.Next() {
		var product ent.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
			return products, pkg.InternalServerError(err, "Mapping struct error")
		}
		products = append(products, product)
	}

	return products, nil
}

func CreateProduct(name string, price float32) *pkg.HttpError {
	_, err := database.GoquDb.
		From(database.ProductTable).
		Insert().
		Rows(goqu.Record{
			"name":  name,
			"price": price,
		}).
		Executor().
		Exec()
	if err != nil {
		return pkg.ConflictError(err, "can't create new product", &map[string][]string{
			"name": {fmt.Sprintf("value '%s' already exist", name)},
		})
	}

	return nil
}

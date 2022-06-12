package database

import (
	"github.com/IlyaSkorychenko/simple_shop_BE/pkg"
	"github.com/doug-martin/goqu/v9"
)

func Seed() {
	seedProducts()
}

func seedProducts() {
	_, err := GoquDb.
		From(ProductTable).
		Insert().
		Rows(
			goqu.Record{"name": "first_product", "price": 199},
			goqu.Record{"name": "second_product", "price": 300},
		).
		Executor().
		Exec()
	pkg.Check(err)
}

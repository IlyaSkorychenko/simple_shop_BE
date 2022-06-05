package database

import (
	"github.com/IlyaSkorychenko/simple_shop_BE/pkg"
	"github.com/doug-martin/goqu/v9"
)

func Seed() {
	seedProducts()
}

func seedProducts() {
	query := GoquDb.
		Insert("products").
		Rows(
			goqu.Record{"name": "first_product", "price": 199},
			goqu.Record{"name": "second_product", "price": 300},
		)
	sql, args, _ := query.ToSQL()
	_, err := GoquDb.Exec(sql, args...)
	pkg.Check(err)
}

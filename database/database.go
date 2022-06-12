package database

import (
	"database/sql"
	"fmt"
	"github.com/IlyaSkorychenko/simple_shop_BE/pkg"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/lib/pq"
)

var GoquDb *goqu.Database
var Client *sql.DB

const (
	ProductTable = "products"
)

func Connect() {
	var (
		host     = pkg.GetEnv("DB_HOST")
		port     = pkg.GetEnv("DB_PORT")
		user     = pkg.GetEnv("DB_USER")
		password = pkg.GetEnv("DB_PASSWORD")
		schema   = pkg.GetEnv("DB_SCHEMA")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, schema)
	client, err := sql.Open("postgres", psqlInfo)
	pkg.Check(err)
	pingErr := client.Ping()
	pkg.Check(pingErr)
	dialect := goqu.Dialect("postgres")
	GoquDb = dialect.DB(client)
	Client = client
}

package main

import (
	"flag"
	"github.com/IlyaSkorychenko/simple_shop_BE/api"
	"github.com/IlyaSkorychenko/simple_shop_BE/database"
	"github.com/IlyaSkorychenko/simple_shop_BE/pkg"
)

func main() {
	flag.Parse()
	pkg.ReadEnv()

	database.Connect()
	defer func() {
		err := database.Client.Close()
		pkg.Check(err)
	}()

	server := api.NewServer()
	server.Start(":" + pkg.GetEnv("PORT"))
}

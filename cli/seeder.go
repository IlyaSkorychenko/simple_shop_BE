package cli

import "github.com/IlyaSkorychenko/simple_shop_BE/database"

const (
	run = "run"
)

func runSeeder(_ string) {
	database.Seed()
}

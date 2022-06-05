package cli

import (
	"github.com/IlyaSkorychenko/simple_shop_BE/database"
)

const (
	up   = "up"
	down = "down"
)

func runMigration(flagValue string) {
	isUp := flagValue == up
	database.Migrate(isUp)
}

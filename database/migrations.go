package database

import (
	"flag"
	"fmt"
	"github.com/IlyaSkorychenko/simple_shop_BE/pkg"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"os"
	"regexp"
)

var migrationPath string

func init() {
	flag.StringVar(&migrationPath, "path", "migrations", "Path to migrations sql files")
}

func Migrate(upOnly bool) {
	migrations := readMigrations(upOnly)

	if upOnly {
		fmt.Print("\n== Execute mode=UP ==\n")
	} else {
		fmt.Print("\n== Execute mode=DOWN ==\n")
	}

	for filename, query := range migrations {
		_, err := GoquDb.Exec(query)
		pkg.Check(err)
		fmt.Printf("\n-- Execudet SQL script \"%s\" --\n", filename)
	}
}

func readMigrations(upOnly bool) map[string]string {
	migrations := make(map[string]string)
	sqlPath := migrationPath + "/"

	dir, err := os.Open(sqlPath)
	pkg.Check(err)
	defer func(file *os.File) {
		err := file.Close()
		pkg.Check(err)
	}(dir)

	filesInfo, err := dir.Readdir(0)
	pkg.Check(err)

	for _, v := range filesInfo {
		match, err := regexp.MatchString("\\w+_(down|up).sql", v.Name())
		pkg.Check(err)

		if !match {
			continue
		}

		filenameLen := len(v.Name())
		isUp := v.Name()[filenameLen-5] == 'p'
		file, err := os.ReadFile(sqlPath + v.Name())
		pkg.Check(err)

		if upOnly != isUp {
			continue
		}

		migrations[v.Name()] = string(file)
	}

	return migrations
}

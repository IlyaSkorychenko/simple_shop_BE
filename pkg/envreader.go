package pkg

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"syscall"
)

var envPath string

func init() {
	flag.StringVar(&envPath, "env_path", ".env", "Path to environment file")
}

func ReadEnv() {
	file, err := os.Open(envPath)
	CheckWithCustomError(err, fmt.Sprintf("Config file doesn't exist. Path: %s", envPath))

	defer func(file *os.File) {
		err := file.Close()
		Check(err)
	}(file)

	scanner := bufio.NewScanner(file)
	var values []string
	var row string

	for scanner.Scan() {
		row = scanner.Text()
		matched, err := regexp.MatchString("^[A-Z_]+=(\\w|\"+.+\")+$", row)
		Check(err)

		if matched {
			values = strings.Split(row, "=")
			values[1] = strings.Trim(values[1], "\"")
			err := os.Setenv(values[0], values[1])
			Check(err)
		}
	}
}

func GetEnv(key string) string {
	value, ok := syscall.Getenv(key)
	if !ok {
		panic(fmt.Sprintf("Value of %s doesn't exist", key))
	}

	return value
}

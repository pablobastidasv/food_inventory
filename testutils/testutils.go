package testutils

import (
	"database/sql"
	"log/slog"
	"os"
	"regexp"
	"sync"

	"github.com/joho/godotenv"
	"github.com/pablobastidasv/fridge_inventory/db"
)

var lock = &sync.Mutex{}

var testDatabase *sql.DB

func DbInstance() *sql.DB {
	if testDatabase == nil {
		lock.Lock()
		defer lock.Unlock()

		if testDatabase == nil {
			testDatabase = db.NewPostgresDb()
		}
	}

	return testDatabase
}

// https://github.com/joho/godotenv/issues/43#issuecomment-503183127
func LoadEnv() {
	re := regexp.MustCompile(`^(.*food_inventory)`) // the name of the project folder MUST be `food_inventory`
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		slog.Error(
			"Problem loading .env file, continuing with environment variables",
			"cause", err,
			"cwd", cwd,
		)
	}
}

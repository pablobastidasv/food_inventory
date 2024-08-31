package db

import (
	"database/sql"
	"os"
)

func NewPostgresDb() *sql.DB {
	connStr := os.Getenv("DBSTRING")
	if connStr == "" {
		panic("DBSTRING environment variable must be provided to initialize the server")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		panic("DATABASE_URL environment variable required but not set")
	}

	db, err := sql.Open("postgres", databaseUrl)

	if err != nil {
		panic(err)
	}

	return db
}

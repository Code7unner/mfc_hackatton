package db

import (
	"database/sql"
	"log"
	"os"
)

func Connect() (*sql.DB, error) {
	url, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Fatalln("$DATABASE_URL is required")
	}

	db, err := sql.Open("postgres", url)
	return db, err

}

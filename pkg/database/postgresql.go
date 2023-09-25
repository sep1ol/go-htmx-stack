package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectPostgres(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func ClosePostgres(db *sql.DB) error {
	return db.Close()
}

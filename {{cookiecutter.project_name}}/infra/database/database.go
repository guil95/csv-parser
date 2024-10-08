package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

func NewPostgresDB() *sqlx.DB {
	dbSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password= %s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)
	db, _ := sqlx.Connect("postgres", dbSource)

	return db
}

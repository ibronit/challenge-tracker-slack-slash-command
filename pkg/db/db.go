package db

import (
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

var db *pg.DB

func init() {
	godotenv.Load()

	dbHost := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	db = pg.Connect(&pg.Options{
		Addr:     dbHost,
		User:     username,
		Password: password,
		Database: dbName,
	})
}

func GetDB() *pg.DB {
	return db
}

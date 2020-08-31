package db

import (
	"os"

	"github.com/ibronit/challenge-tracker-slack-slash-command/internal/challenge"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
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

	createSchema(db)
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*challenge.Challenge)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func GetDB() *pg.DB {
	return db
}

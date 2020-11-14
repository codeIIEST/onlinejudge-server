package models

import (
	"context"
	"os"

	pg "github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

// DBConfigURL  config url of database
var DBConfigURL *pg.DB

// InitDb to initialise database
func InitDb() {
	godotenv.Load()

	DBConfigURL = pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWD"),
		Database: os.Getenv("DB_NAME"),
	})

	//Check if the database is running
	ctx := context.Background()

	if err := DBConfigURL.Ping(ctx); err != nil {
		panic(err)
	}
}

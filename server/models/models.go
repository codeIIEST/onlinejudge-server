package models

import (
	"context"
	"fmt"
	"log"

	pg "github.com/go-pg/pg/v10"
	"github.com/raydwaipayan/onlinejudge-server/config"
)

// DBConfigURL  config url of database
var DBConfigURL *pg.DB

// InitDb to initialise database
func InitDb(conf *config.Config) {
	DBConfigURL = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf(":%s", conf.DbPort),
		User:     conf.DbUser,
		Password: conf.DbPass,
		Database: conf.DbName,
	})

	//Check if the database is running
	ctx := context.Background()

	if err := DBConfigURL.Ping(ctx); err != nil {
		log.Fatal(err)
	}
}

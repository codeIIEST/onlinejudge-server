package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table users...")
		_, err := db.Exec(`CREATE TABLE users (
			id serial PRIMARY KEY,
			first_name varchar(30),
			last_name varchar(30),
			email varchar(50) UNIQUE NOT NULL,
			password varchar NOT NULL,
			handle varchar(30) UNIQUE NOT NULL,
			rating integer DEFAULT 1300,
			createdAt timestamp
		);`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table users...")
		_, err := db.Exec(`DROP TABLE users`)
		return err
	})
}

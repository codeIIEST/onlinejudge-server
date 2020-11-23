package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("CREATING TABLE: users")
		_, userErr := db.Exec(`CREATE TABLE users (
			id serial PRIMARY KEY,
			first_name varchar(30),
			last_name varchar(30),
			email varchar(50) UNIQUE NOT NULL,
			password varchar NOT NULL,
			handle varchar(30) UNIQUE NOT NULL,
			rating integer DEFAULT 1300,
			createdAt timestamp
		);`)
		fmt.Println("CREATING TABLE: contests")
		_, contestsErr := db.Exec(`CREATE TABLE contests (
			id serial PRIMARY KEY,
			name varchar(30),
			type varchar(30),
			phase varchar(50),
			duration_seconds integer,
			start_time_seconds integer,
			relative_time_seconds integer,
			author text[],
			problems text[]			
		);`)
		
		if userErr != nil {
			return userErr
		}
		if contestsErr != nil {
			return contestsErr
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("dropping table users...")
		_, err := db.Exec(`DROP TABLE users`)
		return err
	})
}

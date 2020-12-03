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

		if userErr != nil {
			return userErr
		}

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
			problems text[],
			users text[]			
			);`)

		if contestsErr != nil {
			return contestsErr
		}

		fmt.Println("CREATING TABLE: submissions")

		_, submissionsErr := db.Exec(`CREATE TABLE submissions (
			id serial PRIMARY KEY,
			user_id varchar(50),
			test_result_id varchar(50),
			test_data_id varchar(50),
			contest_id varchar(50),
			problem_id varchar(50),
			creation_time timestamp NOT NULL DEFAULT NOW(),
			relative_time timestamp,
			verdict varchar(30),
			language varchar(10),
			passed_test_count integer,
			time_consumed decimal,
			memory_consumed integer,
			point decimal
		);`)

		if submissionsErr != nil {
			return submissionsErr
		}

		fmt.Println("CREATING TABLE: problems")
		_, problemsErr := db.Exec(`CREATE TABLE problems (
			id serial PRIMARY KEY,
			contest_id varchar(20) NOT NULL,
			index varchar(10) UNIQUE,
			name varchar(30),
			rating integer DEFAULT 800,
			tags text[],
			statement varchar(3000) NOT NULL,
			created_at timestamp NOT NULL DEFAULT NOW()
		);`)

		if problemsErr != nil {
			return problemsErr
		}

		return nil
	}, func(db migrations.DB) error {
		fmt.Println("dropping table users...")
		_, err := db.Exec(`DROP TABLE users`)
		return err
	})
}

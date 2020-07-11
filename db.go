package main

import (
	"database/sql"
)

var db *sql.DB

func connection() error {
	conStr := "user=Postgres password=60nurilla dbname=postgres sslmode=disable"
	_, err := sql.Open("postgres", conStr)
	if err != nil {
		return err

	}
	return nil
}

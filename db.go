package main

import (
	"database/sql"
)

var db *sql.DB

func connection() error {
	var err error
	conStr := "user=postgres password=60nurilla dbname=postgo sslmode=disable"
	db, err = sql.Open("postgres", conStr)
	if err != nil {
		return err

	}
	return nil
}

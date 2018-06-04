package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)


func getDBConnetion() (*sql.DB, error) {
	db, err := sql.Open("postgres", "dbname=lss_db user=postgres password=postgres sslmode=disable")
	return db, err
}
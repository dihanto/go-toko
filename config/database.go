package config

import (
	"database/sql"
	"time"

	"github.com/dihanto/go-toko/helper"
	_ "github.com/lib/pq"
)

func NewDb() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=mastermind sslmode=disable")
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetMaxIdleConns(4)
	db.SetMaxOpenConns(10)

	return db

}

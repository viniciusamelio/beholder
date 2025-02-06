package services

import (
	"database/sql"

	_ "github.com/tursodatabase/go-libsql"
)

func InitSqlite() *sql.DB {
	db, err := sql.Open("libsql", "file:./beholder.db")
	if err != nil {
		panic(err)
	}
	return db
}

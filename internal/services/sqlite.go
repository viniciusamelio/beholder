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
	db.Exec("PRAGMA journal_mode = WAL; PRAGMA synchronous = normal; PRAGMA journal_size_limit = 6144000;")
	return db
}

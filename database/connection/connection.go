package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:7FZijfGP-WwAK!8PwHYTJ7vdGmC9g@/")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	exec(db, "create database if not exists said_db")
	exec(db, "use said_db")
}

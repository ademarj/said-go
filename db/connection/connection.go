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
	exec(db, "drop table if exists contact")
	exec(db, `create table contact(
		id_number varchar(13),
		date_of_birth date NOT NULL,
		gender varchar(50),
		sa_citizen boolean,
		counter int(10),
		PRIMARY KEY (id_number)
	)`)

}

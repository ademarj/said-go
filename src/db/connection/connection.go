package main

import (
	"database/sql"

	"github.com/ademarj/said-go/src/db"
	_ "github.com/go-sql-driver/mysql"
)

func exec(con *sql.DB, sql string) sql.Result {
	result, err := con.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	con := db.Connect()
	defer con.Close()

	exec(con, "create database if not exists said_db")
	exec(con, "use said_db")
	exec(con, "drop table if exists contact")
	exec(con, `create table contact(
		id_number varchar(13),
		date_of_birth date NOT NULL,
		gender varchar(50),
		sa_citizen boolean,
		counter int(10),
		PRIMARY KEY (id_number)
	)`)

}

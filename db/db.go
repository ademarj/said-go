package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const DRIVER = "mysql"
const USER = "root"
const PASS = "7FZijfGP-WwAK!8PwHYTJ7vdGmC9g"
const DBNAME = "said_db"

func Connect() *sql.DB {

	URL := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", USER, PASS, DBNAME)

	con, erro := sql.Open(DRIVER, URL)

	if erro != nil {
		fmt.Println("::::No::::")
		panic(erro.Error())
	}

	return con

}

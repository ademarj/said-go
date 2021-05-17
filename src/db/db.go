package db

import (
	"database/sql"
	"fmt"

	"github.com/ademarj/said-go/src/util/security"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DRIVER                 = "mysql"
	URL_DATABASE_SAID      = "URL_DATABASE_SAID"
	PORT_DATA_BASE         = "PORT_DATA_BASE"
	USER_DATABASE_SAID     = "USER_DATABASE_SAID"
	PASSWORD_DATABASE_SAID = "PASSWORD_DATABASE_SAID"
	NAME_DATABASE          = "NAME_DATABASE"
	URL_CONNECTION         = "%s:%s@tcp(%s:%s)/%s"
)

func Connect() *sql.DB {
	con, erro := sql.Open(DRIVER, makeUrlConnection())
	if erro != nil {
		panic(erro.Error())
	}
	return con
}

func makeUrlConnection() string {
	nameDB, _ := security.GetSecret(NAME_DATABASE)
	urlDB, _ := security.GetSecret(URL_DATABASE_SAID)
	portDB, _ := security.GetSecret(PORT_DATA_BASE)
	userDB, _ := security.GetSecret(USER_DATABASE_SAID)
	passwordDB, _ := security.GetSecret(PASSWORD_DATABASE_SAID)
	return fmt.Sprintf(URL_CONNECTION, userDB, passwordDB, urlDB, portDB, nameDB)
}

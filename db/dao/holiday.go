package dao

import (
	"fmt"

	"github.com/ademarj/said-go/db"
	"github.com/ademarj/said-go/db/model"
	_ "github.com/go-sql-driver/mysql"
)

func GetHolidaysFrom(contactId string) ([]model.Holiday, error) {
	con := db.Connect()
	sql := "select id, name, date, description, id_number from holiday where id_number = ?"
	rs, erro := con.Query(sql, contactId)

	if erro != nil {
		return nil, erro
	}

	var holidays model.Holidays
	for rs.Next() {
		var holiday model.Holiday
		erro := rs.Scan(&holiday.Id, &holiday.Name, &holiday.Date, &holiday.Description, &holiday.ContactId)

		if erro != nil {
			return nil, erro
		}

		holidays = append(holidays, holiday)
	}

	defer rs.Close()
	defer con.Close()

	return holidays, nil
}

func SaveHoliday(holiday model.Holiday) (bool, error) {
	con := db.Connect()
	sql := "insert into holiday (id, name, date, description, id_number) values (?, ?, ?, ?, ?)"
	stmt, erro := con.Prepare(sql)

	if erro != nil {
		return false, erro
	}

	fmt.Printf("\nHoliday Id   %s\n", holiday.Id)
	fmt.Printf("\nHoliday Name %s\n", holiday.Name)
	fmt.Printf("\nHoliday Date %s\n", holiday.Date)
	fmt.Printf("\nDescription  %s\n", holiday.Description)
	fmt.Printf("\nContactId    %s\n", holiday.ContactId)

	_, erro = stmt.Exec(holiday.Id, holiday.Name, holiday.Date, holiday.Description, holiday.ContactId)

	if erro != nil {
		return false, erro
	}

	defer stmt.Close()
	defer con.Close()

	return true, nil
}

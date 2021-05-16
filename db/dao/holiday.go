package dao

import (
	"log"

	"github.com/ademarj/said-go/db"
	"github.com/ademarj/said-go/db/model"
	_ "github.com/go-sql-driver/mysql"
)

func GetHolidaysFrom(contactId string) ([]model.Holiday, error) {
	con := db.Connect()
	sql := "select id, name, date, description, id_number from holiday where id_number = ?"
	rs, erro := con.Query(sql, contactId)
	if erro != nil {
		return []model.Holiday{}, erro
	}
	var holidays model.Holidays
	for rs.Next() {
		var holiday model.Holiday
		erro := rs.Scan(&holiday.Id, &holiday.Name, &holiday.Date, &holiday.Description, &holiday.ContactId)

		if erro != nil {
			return []model.Holiday{}, erro
		}

		holidays = append(holidays, holiday)
	}
	defer rs.Close()
	defer con.Close()

	return holidays, nil
}

func SaveHoliday(holidays []model.Holiday) {
	if len(holidays) > 0 {
		sql := "insert into holiday (id, name, date, description, id_number) values (?, ?, ?, ?, ?)"
		con := db.Connect()
		defer con.Close()
		stmts := []*db.PipelineStmt{}
		for _, holiday := range holidays {
			stmts = append(stmts, db.NewPipelineStmt(sql, holiday.Id, holiday.Name, holiday.Date, holiday.Description, holiday.ContactId))
		}
		err := db.WithTransaction(con, func(tx db.Transaction) error {
			_, err := db.RunPipeline(tx, stmts...)
			return err
		})
		handleError(err)
		log.Println("Done.")
		return
	}

	log.Println("No Holiday to save.")
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

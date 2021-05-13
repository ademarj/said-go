package dao

import (
	"github.com/ademarj/said-go/db"
	"github.com/ademarj/said-go/db/model"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllContacts() ([]model.Contact, error) {
	con := db.Connect()
	sql := "select id_number, date_of_birth, gender, sa_citizen, counter from contact"
	rs, erro := con.Query(sql)

	if erro != nil {
		return nil, erro
	}

	var contacts model.Contacts
	for rs.Next() {
		var contact model.Contact
		erro := rs.Scan(&contact.IdNumber, &contact.DateOfBirthday, &contact.Gender, &contact.SaCitizen, &contact.Counter)

		if erro != nil {
			return nil, erro
		}

		contacts = append(contacts, contact)
	}

	defer rs.Close()
	defer con.Close()

	return contacts, nil
}

func FindBy(idNumber string) (model.Contact, error) {
	con := db.Connect()
	sql := "select id_number, date_of_birth, gender, sa_citizen, counter from contact where id_number = ?"
	rs, erro := con.Query(sql, idNumber)

	if erro != nil {
		return model.Contact{}, nil
	}

	var contact model.Contact

	if rs.Next() {
		erro := rs.Scan(&contact.IdNumber, &contact.DateOfBirthday, &contact.Gender, &contact.SaCitizen, &contact.Counter)

		if erro != nil {
			return model.Contact{}, nil
		}
	}

	defer con.Close()
	defer rs.Close()

	return contact, nil
}

func SaveNewContact(contact model.Contact) (bool, error) {
	con := db.Connect()
	sql := "insert into contact (id_number, date_of_birth, gender, sa_citizen, counter) values (?, ?, ?, ?, ?)"
	stmt, erro := con.Prepare(sql)

	if erro != nil {
		return false, erro
	}

	_, erro = stmt.Exec(contact.IdNumber, contact.DateOfBirthday, contact.Gender, contact.SaCitizen, contact.Counter)

	if erro != nil {
		return false, erro
	}

	defer stmt.Close()
	defer con.Close()

	return true, nil
}

func UpdateContact(contact model.Contact) (bool, error) {
	con := db.Connect()
	sql := "update contact set counter = ? where id_number = ?"
	stmt, erro := con.Prepare(sql)

	if erro != nil {
		return false, erro
	}

	_, erro = stmt.Exec(contact.Counter, contact.IdNumber)

	if erro != nil {
		return false, erro
	}

	defer stmt.Close()
	defer con.Close()

	return true, nil
}

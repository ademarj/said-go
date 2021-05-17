package model

type Holiday struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	Date        string `db:"date"`
	Description string `db:"description"`
	ContactId   string `db:"id_number"`
	Type        string `db:"type"`
}

type Holidays []Holiday

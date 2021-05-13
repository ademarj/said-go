package model

type Holiday struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	Date        string `db:"date"`
	Description string `db:"description"`
	ContactId   string `db:"contact_id_number"`
}

type Holidays struct {
	Holidays []Holiday
}

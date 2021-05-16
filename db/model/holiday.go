package model

type Holiday struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	Date        string `db:"date"`
	Description string `db:"description"`
	ContactId   string `db:"id_number"`
	GridColumn  int
	GridRow     int
}

type Holidays []Holiday

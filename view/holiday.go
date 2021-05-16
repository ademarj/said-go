package view

type Holiday struct {
	Id          string
	Name        string
	Date        string
	Description string
	ContactId   string
	GridColumn  int
	GridRow     int
	Calendar    Calendar
	Type        string
}

type Holidays struct {
	Holidays []Holiday
}

package model

type Contact struct {
	IdNumber        string `db:"id_number"`
	DateOfBirthday  string `db:"date_of_birth"`
	Gender          string `db:"gender"`
	SaCitizen       bool   `db:"sa_citizen"`
	Counter         int32  `db:"counter"`
	DayOfBirthday   int
	MonthOfBirthday int
}

type Contacts []Contact

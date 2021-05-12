package model

import "time"

type Contact struct {
	IdNumber       string    `db:"id_number"`
	DateOfBirthday time.Time `db:"date_of_birth"`
	Gender         string    `db:"gender"`
	SaCitizen      bool      `db:"sa_citizen"`
	Counter        string    `db:"counter"`
}

type Contacts []Contact

package data

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    int
	Nama  string
	Email string
}

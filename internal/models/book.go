package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Name        string
	Author      string
	Publication string
	Description string
	Isbn        int32 `gorm:"unique;not null"`
}

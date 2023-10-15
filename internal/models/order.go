package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	id       int16 `gorm:"unique;not null"`
	Quantity int16
}

package models

import "gorm.io/gorm"

type Tournament struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Date  string `gorm:"not null"`
	Seats []Seat
}

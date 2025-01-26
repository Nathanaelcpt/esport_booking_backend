package models

import "gorm.io/gorm"

type Seat struct {
	gorm.Model
	SeatNumber   string `gorm:"not null"`
	IsBooked     bool   `gorm:"default:false"`
	TournamentID uint
}

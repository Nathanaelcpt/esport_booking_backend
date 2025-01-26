package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	UsersID uint
	SeatID  uint
}

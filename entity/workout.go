package entity

import "time"

type Workout struct {
	ID     uint
	UserID uint
	Date   time.Time
	Notes string
}
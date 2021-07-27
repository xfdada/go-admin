package model

import "time"

type Model struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

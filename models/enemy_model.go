package models

import "gorm.io/gorm"

type Enemy struct {
	gorm.Model
	Healthpoints int `gorm:"default=50"`
	GameID       string
}

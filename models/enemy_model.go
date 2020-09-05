package models

import "gorm.io/gorm"

type Enemy struct {
	gorm.Model
	Healthpoints int
	GameID       string
}

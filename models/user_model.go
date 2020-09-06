package models

import "gorm.io/gorm"

type Creator struct {
	gorm.Model
	Name         string
	Healthpoints int `gorm:"default=30"`
	GameID       string
}

type TeamMate struct {
	gorm.Model
	Name         string
	Healthpoints int `gorm:"default=30"`
	GameID       string
}

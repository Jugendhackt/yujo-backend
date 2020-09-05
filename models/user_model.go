package models

import "gorm.io/gorm"

type Creator struct {
	gorm.Model
	Name         string
	Healthpoints int
	GameID       string
}

type TeamMate struct {
	gorm.Model
	Name         string
	Healthpoints int
	GameID       string
}

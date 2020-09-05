package models

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	GamePin uint64 `json:"game-pin"`
}

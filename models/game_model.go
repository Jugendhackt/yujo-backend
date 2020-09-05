package models

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Game struct {
	ID      string `sql:"primary_key;defautl:uuid_generate_v4()" json:"id"`
	GamePin uint64 `json:"game-pin"`
}

func (game *Game) BeforeCreate(tx *gorm.DB) (err error) {
	game.ID = uuid.New().String()

	if _, uuidErr := uuid.Parse(game.ID); uuidErr != nil {
		err = errors.New("Can't generate valid UUID")
	}
	return
}

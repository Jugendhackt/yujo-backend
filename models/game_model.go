package models

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Game struct {
	ID             string `gorm:"uniqueIndex" sql:"primary_key;defautl:uuid_generate_v4()"`
	GamePin        uint64
	Creator        Creator
	TeamMate       TeamMate
	TeamMateJoined bool
	Enemy          Enemy
}

func (game *Game) BeforeCreate(tx *gorm.DB) (err error) {
	game.ID = uuid.New().String()

	if _, uuidErr := uuid.Parse(game.ID); uuidErr != nil {
		err = errors.New("Can't generate valid UUID")
	}
	return
}

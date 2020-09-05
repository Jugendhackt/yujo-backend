package models

import "gorm.io/gorm"

type Round struct {
	gorm.Model
	GameBaseID int64
	GameID     string
	QuestionID uint
	Answers    []Answer
}

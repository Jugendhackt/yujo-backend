package main

import (
	"bufio"
	"log"

	"github.com/Jugendhackt/yujo-backend/config"
	"github.com/Jugendhackt/yujo-backend/models"
	"github.com/markbates/pkger"
)

func ReadQuestionsFromFile(path string) {
	file, err := pkger.Open(path)
	if err != nil {
		log.Println("Could not open file:", err)
	}

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		var count int64
		config.DB.Table("questions").Where("Text = ?", line).Count(&count)

		if count <= 0 {
			newQuestion := models.Question{
				Text: line,
			}
			config.DB.Save(&newQuestion)
		}
	}
}

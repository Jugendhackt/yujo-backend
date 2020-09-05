package main

import (
	"fmt"
	"log"

	"github.com/Jugendhackt/yujo-backend/config"
	"github.com/Jugendhackt/yujo-backend/models"
	"github.com/Jugendhackt/yujo-backend/router"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Test")

	var err error
	config.DB, err = gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})

	if err != nil {
		log.Println("DB-Status:", err)
	}

	//defer config.DB.Close()

	config.DB.AutoMigrate(&models.Game{})

	r := router.CreateRouter()
	r.Run(":8080")
}

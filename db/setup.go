package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!")

	}

	database.AutoMigrate(&Record{})

	DB = database
}

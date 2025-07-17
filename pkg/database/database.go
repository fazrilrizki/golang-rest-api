package database

import (
	"fmt"
	"golang/golang-rest-api/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Add parseTime=true to the DSN to properly handle time.Time values
	dsn := "root:fazril@tcp(localhost:3306)/belajar-golang?charset=utf8mb4&parseTime=true&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.Article{})

	DB = database
	fmt.Println("Connect to database success")
}

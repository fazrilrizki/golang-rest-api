package database

import (
	"fmt"
	"golang/golang-rest-api/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:fazril@tcp(localhost:3306)/belajar-golang"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.Article{})

	DB = database
	fmt.Println("Connect to database success")
}

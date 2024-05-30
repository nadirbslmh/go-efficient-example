package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Email string `gorm:"type:VARCHAR(255)"`
	Phone string `gorm:"type:VARCHAR(15)"`
	Age   uint8
}

var DB *gorm.DB

func initDB() {
	const (
		DB_USER     = "root"
		DB_PASSWORD = ""
		DB_HOST     = "localhost"
		DB_PORT     = "3306"
		DB_NAME     = "altadb"
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&User{})
}

func createUser() {
	// Insert example data
	user := User{
		Email: "example@example.com",
		Phone: "+1234567890",
		Age:   30,
	}

	result := DB.Create(&user)
	if result.Error != nil {
		log.Println("Error:", result.Error)
	} else {
		fmt.Println("User created:", user)
	}
}

func main() {
	initDB()
	createUser()
}

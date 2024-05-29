package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    uint
	Name  string
	Posts []Post
}

type Post struct {
	ID     uint
	UserID uint
	Title  string
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

	DB.AutoMigrate(&User{}, &Post{})
}

func fetchUsers() {
	var users []User
	DB.Preload("Posts").Find(&users)

	for _, user := range users {
		fmt.Printf("User: %s\n", user.Name)
		for _, post := range user.Posts {
			fmt.Printf("  Post: %s\n", post.Title)
		}
	}
}

func main() {
	initDB()
	fetchUsers()
}

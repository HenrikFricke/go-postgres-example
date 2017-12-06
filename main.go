package main

import (
	_ "github.com/mattn/go-sqlite3"
)

// func InitDb() *gorm.DB {
// 	// Openning file
// 	db, err := gorm.Open("sqlite3", "./data.db")
// 	db.LogMode(true)
// 	// Error
// 	if err != nil {
// 		panic(err)
// 	}
// 	// Creating the table
// 	if !db.HasTable(&models.Users{}) {
// 		db.CreateTable(&models.Users{})
// 		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.Users{})
// 	}

// 	return db
// }

func main() {
	// r := gin.Default()
	// v1 := r.Group("api/v1")

	// v1.POST("/users", PostUser)
	// v1.GET("/users", GetUsers)
	// v1.GET("/users/:id", GetUser)

	// r.Run(":8080")
}

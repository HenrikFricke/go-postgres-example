package main

import (
	"github.com/HenrikFricke/go-postgres-example/api"
	"github.com/HenrikFricke/go-postgres-example/repository"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func initDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./data.db")
	db.LogMode(true)
	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&repository.Users{}) {
		db.CreateTable(&repository.Users{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&repository.Users{})
	}

	return db
}

func main() {
	db := initDb()
	repository := repository.Repository{db}
	api := api.API{&repository}

	r := gin.Default()
	v1 := r.Group("api/v1")

	v1.GET("/users/:id", api.GetUser)

	r.Run(":8080")
}

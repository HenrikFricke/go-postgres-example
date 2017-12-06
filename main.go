package main

import (
	"net/http"

	"github.com/HenrikFricke/go-postgres-example/repository"
	"github.com/HenrikFricke/go-postgres-example/routes"
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
	handler := routes.New(&repository)

	http.ListenAndServe(":8080", handler)
}

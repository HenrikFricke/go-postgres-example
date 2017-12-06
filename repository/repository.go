package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Interface describes the repository
type Interface interface {
	CreateUser(user Users)
	GetUser(id int) (user Users, err error)
	GetUsers() []Users
}

// Repository handles the communication with the database
type Repository struct {
	Db *gorm.DB
}

// CreateUser creates a new user in the database
func (repository *Repository) CreateUser(user Users) {
	repository.Db.Create(&user)
}

// GetUser returns a specific user
func (repository *Repository) GetUser(id int) (user Users, err error) {
	repository.Db.First(&user, id)

	if user.ID == 0 {
		err = errors.New("User not found")
	}

	return user, err
}

// GetUsers returns all users from the database
func (repository *Repository) GetUsers() []Users {
	var users []Users
	repository.Db.Find(&users)

	return users
}

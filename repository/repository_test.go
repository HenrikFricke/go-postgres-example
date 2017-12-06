package repository

import (
	"testing"

	"github.com/dchest/uniuri"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db          *gorm.DB
	transaction *gorm.DB
	repository  Repository
	err         error
)

func init() {
	db, err = gorm.Open("sqlite3", "../test.db")
	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	db.DropTable(&Users{})
	db.CreateTable(&Users{})
}

func BeforeAndAfter() func() {
	transaction = db.Begin()
	repository = Repository{transaction}
	return func() { transaction.Rollback() }
}

// TestCreateUser tests user creation
func TestCreateUser(t *testing.T) {
	defer BeforeAndAfter()()
	user := Users{
		Firstname: uniuri.New(),
		Lastname:  uniuri.New()}

	repository.CreateUser(user)

	var createdUser Users
	transaction.First(&createdUser, 1)

	if createdUser.Firstname != user.Firstname && createdUser.Lastname != createdUser.Lastname {
		t.Error("User was not created.")
	}
}
func TestGetUser(t *testing.T) {
	defer BeforeAndAfter()()
	user := Users{
		Firstname: uniuri.New(),
		Lastname:  uniuri.New()}

	transaction.Create(&user)
	requestedUser, err := repository.GetUser(1)

	if err != nil {
		t.Fatal(err)
	}

	if user.Firstname != requestedUser.Firstname && user.Lastname != requestedUser.Lastname {
		t.Error("GetUser returns wrong user.")
	}
}

func TestGetUsers(t *testing.T) {
	defer BeforeAndAfter()()
	users := []Users{
		Users{
			Firstname: uniuri.New(),
			Lastname:  uniuri.New()},
		Users{
			Firstname: uniuri.New(),
			Lastname:  uniuri.New()},
	}

	for _, user := range users {
		transaction.Create(&user)
	}

	requestedUsers := repository.GetUsers()

	for i, requestedUser := range requestedUsers {
		if requestedUser.Firstname == users[(len(users)-1)-i].Firstname {
			t.Error("GetUsers returns wrong list of users.")
		}
	}
}

package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/HenrikFricke/go-postgres-example/repository"
	"github.com/gin-gonic/gin"
)

var (
	api API
	r   *gin.Engine
	w   *httptest.ResponseRecorder
)

var fakeUser = repository.Users{
	Firstname: "Max",
	Lastname:  "Mustermann"}

type fakeRepository struct{}

func (f *fakeRepository) GetUsers() []repository.Users {
	return []repository.Users{repository.Users{}}
}

func (f *fakeRepository) CreateUser(user repository.Users) {

}

func (f *fakeRepository) GetUser(id int) (user repository.Users, err error) {
	if id == 1 {
		user = fakeUser
	} else {
		err = errors.New("User not found")
	}

	return user, err
}

func Before() {
	api = API{&fakeRepository{}}

	r = gin.Default()
	r.GET("/test/:id", api.GetUser)

	w = httptest.NewRecorder()
}

func TestGetUserSuccessfulRequest(t *testing.T) {
	Before()
	req, _ := http.NewRequest("GET", "/test/1", nil)
	r.ServeHTTP(w, req)

	var parsedBody repository.Users
	err := json.Unmarshal(w.Body.Bytes(), &parsedBody)

	if w.Code != 200 {
		t.Errorf("GetUser handler does not return status code 200, was %s.", strconv.Itoa(w.Code))
	}

	if err != nil || parsedBody.Firstname != fakeUser.Firstname || parsedBody.Lastname != fakeUser.Lastname {
		t.Errorf("GetUser handler does not return proper user in JSON, was %s.", w.Body.String())
	}
}

func TestGetUserUserDoesNotExist(t *testing.T) {
	Before()
	req, _ := http.NewRequest("GET", "/test/2", nil)
	r.ServeHTTP(w, req)

	if w.Code != 404 {
		t.Errorf("GetUser handler does not return status code 404, was %s.", strconv.Itoa(w.Code))
	}
}

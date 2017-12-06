package routes

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/HenrikFricke/go-postgres-example/repository"
	"github.com/stretchr/testify/mock"
)

var (
	res            *httptest.ResponseRecorder
	handler        http.Handler
	mockRepository *MockRepository
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetUsers() []repository.Users {
	args := m.Called()
	return args.Get(0).([]repository.Users)
}

func (m *MockRepository) CreateUser(user repository.Users) {
	m.Called()
}

func (m *MockRepository) GetUser(id int) (user repository.Users, err error) {
	args := m.Called(id)
	return args.Get(0).(repository.Users), args.Error(1)
}

func Before() {
	mockRepository = new(MockRepository)
	handler = New(mockRepository)
	res = httptest.NewRecorder()
}

func TestGetUserSuccessfulRequest(t *testing.T) {
	Before()
	mockRepository.On("GetUser", 1).Return(repository.Users{}, nil)

	req, _ := http.NewRequest("GET", "/v1/users/1", nil)
	handler.ServeHTTP(res, req)

	if res.Code != 200 {
		t.Errorf("GetUser handler does not return status code 200, was %s.", strconv.Itoa(res.Code))
	}

	mockRepository.AssertCalled(t, "GetUser", 1)
}

func TestGetUserUserDoesNotExist(t *testing.T) {
	Before()
	mockRepository.On("GetUser", 2).Return(repository.Users{}, errors.New("Error"))

	req, _ := http.NewRequest("GET", "/v1/users/2", nil)
	handler.ServeHTTP(res, req)

	if res.Code != 400 {
		t.Errorf("GetUser handler does not return status code 404, was %s.", strconv.Itoa(res.Code))
	}

	mockRepository.AssertCalled(t, "GetUser", 2)
}

func TestGetUsers(t *testing.T) {
	Before()
	mockRepository.On("GetUsers").Return([]repository.Users{repository.Users{}, repository.Users{}})

	req, _ := http.NewRequest("GET", "/v1/users", nil)
	handler.ServeHTTP(res, req)

	if res.Code != 200 {
		t.Errorf("GetUsers handler does not return status code 200, was %s.", strconv.Itoa(res.Code))
	}

	mockRepository.AssertCalled(t, "GetUsers")
}

func TestCreateUser(t *testing.T) {
	Before()
	user := repository.Users{}
	mockRepository.On("CreateUser").Return()

	body, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/v1/users", bytes.NewReader(body))
	handler.ServeHTTP(res, req)

	if res.Code != 200 {
		t.Errorf("CreateUser handler does not return status code 200, was %s.", strconv.Itoa(res.Code))
	}

	mockRepository.AssertCalled(t, "CreateUser")
}

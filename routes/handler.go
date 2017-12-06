package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/HenrikFricke/go-postgres-example/repository"
	"github.com/gorilla/mux"
)

// Handler stores all handler for the API
type Handler struct {
	Repository repository.Interface
}

type errorResponse struct {
	Error string `json:"error"`
}

type messageResponse struct {
	Message string `json:"message"`
}

// GetUser returns a specific user
func (h Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user, err := h.Repository.GetUser(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse{err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// GetUsers returns a list of user
func (h Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := h.Repository.GetUsers()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// CreateUser creates a new user
func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user repository.Users
	err := decoder.Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse{err.Error()})
		return
	}

	h.Repository.CreateUser(user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messageResponse{"User has been created."})
}

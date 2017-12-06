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

// GetUser handles request to get all users
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

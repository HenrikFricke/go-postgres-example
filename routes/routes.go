package routes

import (
	"net/http"

	"github.com/HenrikFricke/go-postgres-example/repository"
	"github.com/gorilla/mux"
)

// New returns router for API
func New(repository repository.Interface) http.Handler {
	h := Handler{repository}

	r := mux.NewRouter()
	r.HandleFunc("/users/{id:[0-9]+}", h.GetUser)

	return r
}

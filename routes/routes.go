package routes

import (
	"net/http"

	"github.com/HenrikFricke/go-postgres-example/repository"
	"github.com/gorilla/mux"
)

func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fn(w, r)
	}
}

// New returns router for API
func New(repository repository.Interface) http.Handler {
	h := Handler{repository}

	r := mux.NewRouter()
	r.HandleFunc("/v1/users", addDefaultHeaders(h.GetUsers)).Methods("GET")
	r.HandleFunc("/v1/users", addDefaultHeaders(h.CreateUser)).Methods("POST")
	r.HandleFunc("/v1/users/{id:[0-9]+}", addDefaultHeaders(h.GetUser)).Methods("GET")

	return r
}

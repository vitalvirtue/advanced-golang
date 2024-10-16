package routers

import (
	"github.com/gorilla/mux"
	"github.com/vitalvirtue/Go-Basics/Build-A-CRUD-API/handlers"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", handlers.GetMovie).Methods("GET")
	r.HandleFunc("/movies", handlers.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", handlers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", handlers.DeleteMovie).Methods("DELETE")

	return r
}
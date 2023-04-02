package main

import (
	"log"
	"net/http"

	"github.com/jeypc/go-jwt-mux/controllers/authcontroller"
	"github.com/jeypc/go-jwt-mux/controllers/photocontroller"
	"github.com/jeypc/go-jwt-mux/middlewares"

	"github.com/jeypc/go-jwt-mux/models"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/users/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/users/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/users/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("").Subrouter()
	api.HandleFunc("/photos", photocontroller.Index).Methods("GET")
	api.HandleFunc("/photos/upload", photocontroller.Upload).Methods("POST")
	api.HandleFunc("/photos/{id}", photocontroller.Delete).Methods("DELETE")
	api.HandleFunc("/photos/{id}", photocontroller.Edit).Methods("PUT")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}

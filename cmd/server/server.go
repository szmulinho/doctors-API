package server

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/szmulinho/doctors/internal/api/endpoints/users/get"
	"github.com/szmulinho/doctors/internal/api/endpoints/users/login"
	"github.com/szmulinho/doctors/internal/api/endpoints/users/register"
	"github.com/szmulinho/doctors/internal/api/endpoints/users/userData"
	"log"
	"net/http"
)

func Run() {
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("Starting the application...")
	router.HandleFunc("/login", login.Login).Methods("POST")
	router.HandleFunc("/register", register.CreateUser).Methods("POST")
	router.HandleFunc("/user", userData.GetUserDataHandler).Methods("GET")
	router.HandleFunc("/users", get.GetAllUsers).Methods("GET")
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}),
		handlers.ExposedHeaders([]string{}),
		handlers.AllowCredentials(),
		handlers.MaxAge(86400),
	)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8082"), cors(router)))
}

func server() {
	Run()
}

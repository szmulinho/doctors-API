package server

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/szmulinho/doctors/internal/api/endpoints/doctors/get"
	"github.com/szmulinho/doctors/internal/api/endpoints/doctors/login"
	"github.com/szmulinho/doctors/internal/api/endpoints/doctors/register"
	"github.com/szmulinho/doctors/internal/api/endpoints/doctors/userData"
	"github.com/szmulinho/doctors/internal/api/jwt"
	"log"
	"net/http"
)

func Run() {
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("Starting the application...")
	router.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		userID := uint(1)
		token, err := jwt.GenerateToken(w, r, int64(userID), true)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(token))
	}).Methods("POST")
	router.HandleFunc("/login", login.Login).Methods("POST")
	router.HandleFunc("/register", register.CreateDoctor).Methods("POST")
	router.HandleFunc("/user", doctorData.GetDoctorDataHandler).Methods("GET")
	router.HandleFunc("/doctors", get.GetAllDoctor).Methods("GET")
	router.HandleFunc("/doctor", get.GetDoctorDataHandler).Methods("GET")
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Content-Type"}),
		handlers.ExposedHeaders([]string{}),
		handlers.AllowCredentials(),
		handlers.MaxAge(86400),
	)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8085"), cors(router)))
}

func server() {
	Run()
}

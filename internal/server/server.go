package server

import (
	"context"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/szmulinho/doctors/internal/server/endpoints"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Run(ctx context.Context, db *gorm.DB) {
	handler := endpoints.NewHandler(db)
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("Starting the application...")
	router.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		userID := uint(1)
		token, err := handler.GenerateToken(w, r, int64(userID), true)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(token))
	}).Methods("POST")
	router.HandleFunc("/login", handler.Login).Methods("POST")
	router.HandleFunc("/register", handler.RegisterDoctor).Methods("POST")
	router.HandleFunc("/user", handler.GetDoctorDataHandler).Methods("GET")
	router.HandleFunc("/doctors", handler.GetAllDoctors).Methods("GET")
	router.HandleFunc("/doctor", handler.GetDoctorDataHandler).Methods("GET")
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Content-Type"}),
		handlers.ExposedHeaders([]string{}),
		handlers.AllowCredentials(),
		handlers.MaxAge(86400),
	)
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", "https://szmul-med-doctors.onrender.com/"), cors(router))
		if err != nil {
			log.Fatal(err)
		}
	}()
	<-ctx.Done()
}

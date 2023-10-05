package register

import (
	"encoding/json"
	"net/http"

	"github.com/szmulinho/doctors/internal/database"
	"github.com/szmulinho/doctors/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func CreateDoctor(w http.ResponseWriter, r *http.Request) {
	var newDoctor model.Doctor

	err := json.NewDecoder(r.Body).Decode(&newDoctor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newDoctor.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newDoctor.Password = string(hashedPassword)

	newDoctor.Role = "doctor"

	result := database.DB.Create(&newDoctor)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	doctorJSON, err := json.Marshal(newDoctor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(doctorJSON)
}

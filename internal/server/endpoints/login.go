package endpoints

import (
	"encoding/json"
	"github.com/szmulinho/doctors/internal/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (h *handlers) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var doctor model.Doctor
	result := h.db.Where("login = ?", credentials.Login).First(&doctor)
	if result.Error != nil {
		http.Error(w, "Invalid login or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(doctor.Password), []byte(credentials.Password))
	if err != nil {
		http.Error(w, "Invalid login or password", http.StatusUnauthorized)
		return
	}

	var isDoctor bool

	if doctor.Role == "doctor" {
		isDoctor = true
	}

	token, err := h.GenerateToken(w, r, doctor.ID, isDoctor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := model.LoginResponse{
		Doctor: doctor,
		Token:  token,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

package endpoints

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/szmulinho/doctors/internal/model"
	"log"
	"net/http"
	"strings"
)

func (h *handlers) GetDoctorDataHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	doctor, err := h.GetDoctorFromToken(token)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(doctor)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(response)
}

func (h *handlers) GetDoctorFromToken(tokenString string) (*model.Doctor, error) {
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return model.JwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Invalid token claims")
	}

	doctorID := int64(claims["doctorID"].(float64))

	var doctor model.Doctor
	if err := h.db.First(&doctor, doctorID).Error; err != nil {
		fmt.Printf("Error retrieving doctor from the database: %v\n", err)
		fmt.Printf("Doctor ID: %d\n", doctorID)
		return nil, err
	}

	fmt.Printf("Doctor retrieved from the database: %+v\n", doctor)

	return &doctor, nil
}

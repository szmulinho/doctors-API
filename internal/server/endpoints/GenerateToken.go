package endpoints

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/szmulinho/doctors/internal/model"
	"net/http"
	"time"
)

func (h *handlers) GenerateToken(w http.ResponseWriter, r *http.Request, ID int64, isDoctor bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"doctorID": ID,
		"isDoctor": isDoctor,
		"exp":      time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	})

	tokenString, err := token.SignedString(model.JwtKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", err
	} else {
		fmt.Println("token generated")
	}

	return tokenString, nil
}

package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/szmulinho/doctors/internal/model"
	"net/http"
	"time"
)

func CreateToken(w http.ResponseWriter, r *http.Request) {
	var user model.JwtUser
	_ = json.NewDecoder(r.Body).Decode(&user)
	role := "user"
	if user.Role == "doctor" {
		role = "doctor"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Jwt,
		"password": user.Password,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	})
	tokenString, error := token.SignedString(model.JwtKey)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Token created")
	}
	json.NewEncoder(w).Encode(model.JwtToken{Token: tokenString})
}

package model

import (
	"os"
)

var JwtKey = []byte(os.Getenv("JWT_KEY"))

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

type Response struct {
	Data string `json:"data"`
}

type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	Login    string `gorm:"unique" json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type JwtUser struct {
	Jwt      string "jwt"
	Password string "password"
	Role     string `json:"role"`
}

type LoginResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

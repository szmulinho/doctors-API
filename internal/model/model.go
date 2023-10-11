package model

import (
	"os"
)

var JwtKey = []byte(os.Getenv("JWT_KEY"))

type Exception struct {
	Message string `json:"message"`
}

type Doctor struct {
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	Login    string `gorm:"unique" json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginResponse struct {
	Doctor Doctor `json:"doctor"`
	Token  string `json:"token"`
}

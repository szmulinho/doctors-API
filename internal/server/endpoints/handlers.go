package endpoints

import (
	"github.com/szmulinho/doctors/internal/model"
	"gorm.io/gorm"
	"net/http"
)

type Handlers interface {
	GetAllDoctors(w http.ResponseWriter, r *http.Request)
	GetDoctorDataHandler(w http.ResponseWriter, r *http.Request)
	GetDoctorFromToken(tokenString string) (*model.Doctor, error)
	Login(w http.ResponseWriter, r *http.Request)
	RegisterDoctor(w http.ResponseWriter, r *http.Request)
	GenerateToken(w http.ResponseWriter, r *http.Request, ID int64, isDoctor bool) (string, error)
	ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc
}

type handlers struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) Handlers {
	return &handlers{
		db: db,
	}
}

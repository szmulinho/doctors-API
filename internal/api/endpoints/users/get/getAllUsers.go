package get

import (
	"encoding/json"
	"github.com/szmulinho/doctors/internal/database"
	"github.com/szmulinho/doctors/internal/model"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var Users []model.User
	if err := database.DB.Find(&Users).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(Users)
}

package get

import (
	"encoding/json"
	"github.com/szmulinho/doctors/internal/database"
	"github.com/szmulinho/doctors/internal/model"
	"net/http"
)

func GetAllDoctor(w http.ResponseWriter, r *http.Request) {
	var Doctors []model.Doctor
	if err := database.DB.Find(&Doctors).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(Doctors)
}

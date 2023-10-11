package endpoints

import (
	"encoding/json"
	"github.com/szmulinho/doctors/internal/model"
	"net/http"
)

func (h *handlers) GetAllDoctors(w http.ResponseWriter, r *http.Request) {
	var Doctors []model.Doctor
	if err := h.db.Find(&Doctors).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(Doctors)
}

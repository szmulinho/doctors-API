package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/szmulinho/doctors/internal/model"
)

func TestGetDoctorDataHandler(t *testing.T) {
	err := godotenv.Load(".env.test")
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")

	dsn := "host=" + host + " user=" + user + " dbname=" + name + " sslmode=require password=" + password + " port=" + port

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	defer db.DB()

	db.AutoMigrate(&model.Doctor{})

	h := &handlers{db: db}

	mockDoctor := &model.Doctor{
		ID:   1,
		Login: "doctor",
		Role: "doctor",
	}

	token, err := h.GenerateToken(httptest.NewRecorder(), httptest.NewRequest("GET", "/doctor", nil), mockDoctor.ID, true)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/doctor", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	fmt.Println("Before calling GetDoctorFromToken, db is:", h.db)

	h.GetDoctorDataHandler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseDoctor model.Doctor
	err = json.Unmarshal(w.Body.Bytes(), &responseDoctor)
	assert.NoError(t, err)

	assert.Equal(t, mockDoctor.ID, responseDoctor.ID)
	assert.Equal(t, mockDoctor.Login, responseDoctor.Login)
}

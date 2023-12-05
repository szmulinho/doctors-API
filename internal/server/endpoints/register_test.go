package endpoints

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/szmulinho/doctors/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterDoctor(t *testing.T) {
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

	newDoctor := model.Doctor{
		Login:     "Szmulik",
		Password: "secretpassword",
		Role: "doctor",
	}

	newDoctorJSON, err := json.Marshal(newDoctor)
	assert.NoError(t, err, "Error marshaling new doctor data")

	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(newDoctorJSON))
	assert.NoError(t, err, "Error creating request")

	w := httptest.NewRecorder()

	h.RegisterDoctor(w, req)

	assert.Equal(t, http.StatusCreated, w.Code, "Expected HTTP status Created")

	var responseDoctor model.Doctor
	err = json.Unmarshal(w.Body.Bytes(), &responseDoctor)
	assert.NoError(t, err, "Error decoding response body")

	assert.Equal(t, newDoctor.Login, responseDoctor.Login, "Unexpected doctor name")

}

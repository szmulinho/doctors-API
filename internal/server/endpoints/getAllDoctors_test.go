package endpoints

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/szmulinho/doctors/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetAllDoctors(t *testing.T) {
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

	req, err := http.NewRequest("GET", "/doctors", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	h.GetAllDoctors(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var doctors []model.Doctor
	err = json.Unmarshal(w.Body.Bytes(), &doctors)
	assert.NoError(t, err)

}

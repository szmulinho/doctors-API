package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/szmulinho/common/config"
	"github.com/szmulinho/doctors/internal/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadConfigFromEnv() config.StorageConfig {
	return config.StorageConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("DB_PASSWORD:", os.Getenv("DB_PASSWORD"))
	fmt.Println("DB_NAME:", os.Getenv("DB_NAME"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
}
func Connect() (*gorm.DB, error) {
	config := LoadConfigFromEnv()
	connectionString := config.ConnectionString()

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&model.Doctor{},
	); err != nil {
		return nil, err
	}

	return db, nil
}

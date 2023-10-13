package database

import (
	"github.com/szmulinho/common/config"
	"github.com/szmulinho/common/model"
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

func Connect() (*gorm.DB, error) {
	conn := LoadConfigFromEnv()
	connectionString := conn.ConnectionString()

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate models
	if err := db.AutoMigrate(&model.Prescription{}, &model.Drug{}, &model.User{}, &model.Opinion{}, &model.Order{}, &model.Doctor{}); err != nil {
		return nil, err
	}

	return db, nil
}

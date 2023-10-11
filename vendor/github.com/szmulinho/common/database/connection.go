package database

import (
	"github.com/szmulinho/common/internal/config"
	"github.com/szmulinho/common/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	connectionString := config.StorageConfig{}.ConnectionString()
	conn := postgres.Open(connectionString)
	db, err := gorm.Open(conn, &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Doctor{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

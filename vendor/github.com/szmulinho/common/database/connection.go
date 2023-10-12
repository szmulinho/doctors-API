package database

import (
	"github.com/szmulinho/common/config"
	"github.com/szmulinho/common/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	connectionString := config.StorageConfig{}.ConnectionString()

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&model.Doctor{},
		&model.Opinion{},
		&model.User{},
		&model.Prescription{},
		&model.Drug{},
		&model.Order{},
	); err != nil {
		return nil, err
	}

	return db, nil
}

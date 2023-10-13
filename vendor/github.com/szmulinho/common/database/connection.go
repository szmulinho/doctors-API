package database

import (
	"github.com/szmulinho/common/config"
	"github.com/szmulinho/common/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	conn := config.LoadConfigFromEnv()
	connectionString := conn.ConnectionString()

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&model.Prescription{}, &model.Drug{}, &model.User{}, &model.Opinion{}, &model.Order{}, &model.Doctor{}); err != nil {
		return nil, err
	}

	return db, nil
}

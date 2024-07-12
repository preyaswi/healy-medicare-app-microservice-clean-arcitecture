package db

import (
	"doctor-service/pkg/config"
	"doctor-service/pkg/domain"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	db.AutoMigrate(domain.Doctor{})
	db.AutoMigrate(domain.Review{})
	return db, dbErr
}

package db

import (
	"fmt"
	"healy-admin/pkg/config"
	"healy-admin/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(&domain.Admin{})
	db.AutoMigrate(&domain.Booking{})
	db.AutoMigrate(&domain.RazerPay{})
	db.AutoMigrate(&domain.Prescription{})
	db.AutoMigrate(&domain.Availability{})
	db.AutoMigrate(&domain.Event{})
	return db, dbErr

}

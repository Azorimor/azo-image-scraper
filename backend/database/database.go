package database

import (
	"azorimor/azo-image-scraper/models"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConn *gorm.DB
var sqliteDBName string = "azo-image-scraper.sqlite3"

func CreateDBConnection() error {
	db, err := gorm.Open(sqlite.Open(sqliteDBName), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	dbConn = db

	return err
}

func GetDBConnection() (*gorm.DB, error) {
	sqlDB, err := dbConn.DB()
	if err != nil {
		return dbConn, err
	}
	if err := sqlDB.Ping(); err != nil {
		return dbConn, err
	}
	return dbConn, nil
}

func AutoMigrateDB() error {
	db, connErr := GetDBConnection()
	if connErr != nil {
		return connErr
	}
	// Add required models here
	err := db.AutoMigrate(&models.Image{})
	// Example for migrating multiple models
	// err:= db.AutoMigrate(&models.User{}, &models.Admin{}, &models.Guest{})
	return err
}

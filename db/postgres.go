package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func initDB(config Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Password, config.DBName, config.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func StartDbConnection(config Config) {
	database = initDB(config)
}

func GetDBConn() *gorm.DB {
	return database
}

func CloseDbConnection() error {
	db, err := database.DB()

	if err != nil {
		return fmt.Errorf("error occured on database connection closing: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		return fmt.Errorf("error occured on database connection closing: %s", err.Error())
	}
	return nil
}

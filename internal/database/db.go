package database

import (
	"davifiber/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func getConfig() string {

	dbHost := config.Config("DB_HOST")
	dbName := config.Config("DB_NAME")
	dbUser := config.Config("DB_USER")
	dbPasswd := config.Config("DB_PASSWORD")
	dbPort := config.Config("DB_PORT")
	dbSSLMode := config.Config("DB_SSL_MODE")

	dburi := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", dbHost, dbUser, dbPasswd, dbName, dbPort, dbSSLMode)
	return dburi

}

func Connect() error {
	var err error
	uriConn := getConfig()

	Conn, err = gorm.Open(postgres.Open(uriConn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil

}

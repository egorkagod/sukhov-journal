package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBCredentials struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int64
}

func NewDB(credentials *DBCredentials) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", credentials.Host, credentials.User, credentials.Password, credentials.Name, credentials.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

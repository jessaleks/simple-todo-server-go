package utils

import (
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

var DB, _ = ConnectToTheDatabase()

func ConnectToTheDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil

}

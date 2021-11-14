package utils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToTheDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

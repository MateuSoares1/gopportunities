package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	//initialize sqlite

	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("Error initialize sqlite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	// initialize logger
	logger = NewLogger(prefix)
	return logger
}

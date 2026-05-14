package config

import (
	"os"

	"github.com/Erick-VP/godin/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("sqlite db not found, creating...")

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("failed to create db directory %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("failed to create db file %v", err)
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open("./db/main.db"), &gorm.Config{})

	if err != nil {
		logger.Errorf("failed to initialize sqlite %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("failed to auto migrate opening %v", err)
		return nil, err
	}

	return db, nil
}

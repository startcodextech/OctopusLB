package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

type Config struct {
	db *gorm.DB
}

func Init() (*Config, error) {
	log := logger.Default.LogMode(logger.Warn)

	if os.Getenv("APP_ENV") == "" {
		log = logger.Default.LogMode(logger.Silent)
	}

	db, err := gorm.Open(sqlite.Open("octopus.config"), &gorm.Config{
		Logger: log,
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Module{}, LbApp{}, Certificates{})
	if err != nil {
		return nil, err
	}

	cfg := &Config{db: db}

	return cfg, nil
}

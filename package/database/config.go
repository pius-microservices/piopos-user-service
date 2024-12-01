package database

import (
	"fmt"
	envConfig "github.com/pius-microservices/piopos-user-service/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	envCfg := envConfig.LoadConfig()

	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s ", envCfg.UserDBHost, envCfg.UserDBPort, envCfg.UserDBUsername, envCfg.UserDBPassword, envCfg.UserDBName)

	gormDb, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return gormDb, nil
}
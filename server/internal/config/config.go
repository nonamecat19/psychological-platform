package config

import (
	"os"
	"strconv"
)

type DbConfig struct {
	Database string
	Password string
	Username string
	Port     string
	Host     string
}

func GetDbConfig() DbConfig {
	return DbConfig{
		Database: os.Getenv("DB_DATABASE"),
		Password: os.Getenv("DB_PASSWORD"),
		Username: os.Getenv("DB_USERNAME"),
		Port:     os.Getenv("DB_PORT"),
		Host:     os.Getenv("DB_HOST"),
	}
}

type AppConfig struct {
	Port int
}

func GetAppConfig() AppConfig {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	return AppConfig{
		Port: port,
	}
}

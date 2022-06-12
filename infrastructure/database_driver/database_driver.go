package database_driver

import (
	"github.com/joho/godotenv"
)

type DriverDatabase struct {
	Host     string
	Password string
	Port     string
	User     string
	Name     string
}

func readEnv() (DriverDatabase, error) {
	envDb, err := godotenv.Read()
	if err != nil {
		return DriverDatabase{}, err
	}
	return DriverDatabase{
		Host:     envDb["DB_HOST"],
		Password: envDb["DB_PASSWORD"],
		Port:     envDb["DB_PORT"],
		User:     envDb["DB_USER"],
		Name:     envDb["DB_NAME"],
	}, nil
}

func NewDriverDatabase() (DriverDatabase, error) {
	dataEnv, err := readEnv()
	if err != nil {
		return DriverDatabase{}, err
	}
	return dataEnv, nil
}

package config

import (
	"BE_WEB_BEM_Proker/model"
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dB  *gorm.DB
	err error
)

func InitDB() (*gorm.DB, error) {
	env, errorEnv := godotenv.Read()
	if errorEnv != nil {
		panic("ENV is not set")
	}
	// GEt env
	host := env["DB_HOST"]
	password := env["DB_PASSWORD"]
	port := env["DB_PORT"]
	user := env["DB_USER"]
	name := env["DB_NAME"]
	dsnXAMPP := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, name)
	dB, err = gorm.Open(mysql.Open(dsnXAMPP), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := dB.AutoMigrate(&model.EntitasBPH{}, &model.EntitasProker{}); err != nil {
		panic(err)
	}
	return dB, nil
}

package database_connection

import (
	"BE_WEB_BEM_Proker/domain"
	"BE_WEB_BEM_Proker/infrastructure/database_driver"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MakeConnection(data database_driver.DriverDatabase) (*gorm.DB, error) {
	dsnXAMPP := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", data.User, data.Password, data.Host, data.Port, data.Name)
	db, err := gorm.Open(mysql.Open(dsnXAMPP), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&domain.EntitasProker{}); err != nil {
		return nil, err
	}
	return db, nil
}

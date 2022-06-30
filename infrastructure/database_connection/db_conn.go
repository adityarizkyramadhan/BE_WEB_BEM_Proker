package database_connection

import (
	"BE_WEB_BEM_Proker/domain"
	"BE_WEB_BEM_Proker/infrastructure/database_driver"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MakeConnection(data database_driver.DriverSupabase) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s "+
		"password=%s "+
		"host=%s "+
		"TimeZone=Asia/Singapore "+
		"port=%s "+
		"dbname=%s", data.User, data.Password, data.Host, data.Port, data.DbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&domain.EntitasProker{}, &domain.LinkImage{}, &domain.Admin{}); err != nil {
		return nil, err
	}
	return db, nil
}

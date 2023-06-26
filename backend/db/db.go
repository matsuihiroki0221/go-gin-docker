package db

import (
	"fmt"

	"github.com/matsuihiroki0221/gin-docker/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() error {
	dsn := fmt.Sprintf(settings.Cfg.DB.Dsn,
		settings.Cfg.DB.UserName,
		settings.Cfg.DB.Password,
		settings.Cfg.DB.HostName,
		settings.Cfg.DB.Port,
		settings.Cfg.DB.DBName,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}

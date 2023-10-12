package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"zhongxite/t-template/models"
)

func InitDB() (db *gorm.DB, err error) {
	name := ConfigMysql["name"]
	password := ConfigMysql["password"]
	url := ConfigMysql["url"]
	port := ConfigMysql["port"]
	dbName := ConfigMysql["dbName"]
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", name, password, url, port, dbName)

	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		}),
	})
	CreateInitDB(db)
	return
}

func CreateInitDB(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Menus{})
}

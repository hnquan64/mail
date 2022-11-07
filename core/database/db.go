package database

import (
	"fmt"

	"gitlab.com/meta-node/mail/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbConnection struct {
	DB *gorm.DB
}

var dbConn *dbConnection

func InstanceDB() *dbConnection {
	if dbConn == nil {
		USER := config.GetConfig().DataBaseConfig.User
		PASS := config.GetConfig().DataBaseConfig.Pass
		HOST := config.GetConfig().DataBaseConfig.Host
		PORT := config.GetConfig().DataBaseConfig.Port
		DBNAME := config.GetConfig().DataBaseConfig.DB
		URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
		conn, err := gorm.Open(mysql.Open(URL), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		})
		if err != nil {
			panic(err.Error())
		}
		dbConn = &dbConnection{DB: conn}
	}
	return dbConn
}

package mdb

import (
	"fmt"

	"github.com/lingzerol/simtrans/model/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *gorm.DB
)

func InitDatabase() {
	serverConfig := config.GetServerConfig()
	var err error
	if serverConfig.DBConfig.DBType == "mysql" {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			serverConfig.DBConfig.UserName,
			serverConfig.DBConfig.PassWord,
			serverConfig.DBConfig.Host,
			serverConfig.DBConfig.Port,
			serverConfig.DBConfig.DBName,
		)

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	} else {
		db, err = gorm.Open(sqlite.Open(serverConfig.DBConfig.DBName+".db"),
			&gorm.Config{})
	}
	if err != nil {
		panic("init mysql failed")
	}
}

func GetDBInstance() *gorm.DB {
	return db
}

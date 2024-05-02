package mysql

import "gorm.io/gorm"

var (
	
)

func InitMysql() {
	serverConfig := config.GetServerConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", 
		serverConfig.DBConfig.UserName, serverConfig.DBConfig.PassWord, 
		serverConfig.DBConfig.Host, serverConfig.DBConfig.Port, 
		serverConfig.DBConfig.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

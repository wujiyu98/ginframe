package database

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	dsn := viper.GetString("database.dsn")
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn, // DSN data source name
		DefaultStringSize: 255, // string 类型字段的默认长度
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}

package config

import (
	"backend/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile("config.yaml") // 指定配置文件路径
	if err := viper.ReadInConfig(); err != nil {
		panic("Failed to read the config file: " + err.Error())
	}
}

func ConnectDatabase() *gorm.DB {
	dsn := viper.GetString("database_dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		panic("Failed to connect to database")
	}

	if err = db.AutoMigrate(&model.User{}); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	return db
}

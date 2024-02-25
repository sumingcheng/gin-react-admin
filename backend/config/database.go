package config

import (
	"backend/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic("Failed to read the config file: " + err.Error())
	}
}

func ConnectDatabase() *gorm.DB {
	// 使用viper获取配置项并拼接DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.name"),
		viper.GetString("database.charset"),
		viper.GetString("database.parseTime"),
		viper.GetString("database.loc"),
	)

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

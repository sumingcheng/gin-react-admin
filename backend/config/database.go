package config

import (
	"backend/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := "root:GoGin123!@tcp(127.0.0.1:3306)/adminDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database")
		panic("Failed to connect to database")
	}

	err = db.AutoMigrate(&model.User{})

	return db
}

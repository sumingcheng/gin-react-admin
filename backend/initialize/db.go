package initialize

import (
	"backend/model"
	"fmt"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var once sync.Once
var dbErr error

// 初始化数据库连接
func initDB() {
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

	db, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr == nil {
		dbErr = db.AutoMigrate(&model.User{})
	}
}

// GetDB 提供全局数据库连接
func GetDB() (*gorm.DB, error) {
	once.Do(initDB)
	return db, dbErr
}

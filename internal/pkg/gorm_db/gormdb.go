package gormdb

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (db *gorm.DB, err error) {
	dsn := "root:123456@tcp(127.0.0.1)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	conf := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	db, err = gorm.Open(mysql.Open(dsn), conf)
	if err != nil {
		return
	}

	return
}

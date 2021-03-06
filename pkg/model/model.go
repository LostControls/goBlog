package model

import (
	"goblog/pkg/logger"

	"gorm.io/gorm"
	// GORM 的 MYSQL 驱动的导入
	"gorm.io/driver/mysql"
)

// DB gorm.DB 对象
var DB *gorm.DB

// ConnectDB 初始化模型
func ConnectDB() *gorm.DB {
	var err error

	config := mysql.New(mysql.Config{
		DSN: "root:root@tcp(127.0.0.1:3306)/goblog?charset=utf8&parseTimeTrue&loc=Local",
	})

	// 准备数据库连接池
	DB, err = gorm.Open(config, &gorm.Config{})

	logger.LogError(err)

	return DB
}


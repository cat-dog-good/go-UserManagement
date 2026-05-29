package config

import (
	"fmt"
	"log"

	"go-UserManagement/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	username := "root"
	password := "123456"
	host := "host.docker.internal"
	port := "3307"
	dbname := "go_user_management"

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbname,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败: ", err)
	}

	DB = db
	log.Println("数据库连接成功")

	if err := DB.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("自动迁移失败: ", err)
	}

	log.Println("users 表迁移成功")
}
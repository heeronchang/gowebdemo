package main

import (
	"gowebdemo/internal/pkg/gorm_gen/dao"
	"gowebdemo/internal/pkg/gorm_gen/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "root:123456@tcp(127.0.0.1:3308)/go_web_demo?charset=utf8mb4&parseTime=True&loc=Local"

func InitDB() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	dao.SetDefault(db)
}

func main() {
	u := &model.User{
		Username: "GT",
		Age:      30,
		Password: "123456",
		Phone:    "18510241024",
	}
	err := dao.User.Create(u)
	if err != nil {
		log.Print(err)
		log.Println("11")
	}
}

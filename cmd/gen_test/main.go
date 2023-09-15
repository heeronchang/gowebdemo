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
	InitDB()

	// addOneUser()
	// addMultiUsers()
	// updateOneFiled()
	// updateMultiFields()
	// deleteOne()
	// selectUsers()
	// selectUserPage()
	transaction()
}

func addOneUser() {
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

func addMultiUsers() {
	// 批量插入
	users := []*model.User{
		{
			Username: "小昭",
			Age:      16,
			Password: "zzz369",
			Phone:    "14789632541",
		},
		{
			Username: "韦一笑",
			Age:      18,
			Password: "wyx369",
			Phone:    "14789632542",
		},
		{
			Username: "灭绝师太",
			Age:      50,
			Password: "mjs369",
			Phone:    "14789632543",
		},
	}
	err := dao.User.CreateInBatches(users, 10)
	if err != nil {
		log.Println(err)
	}
}

func updateOneFiled() {
	u := dao.User
	// Update with conditions
	// u.Where(u.ID.Eq(2)).Update(u.Age, 13)

	// Update with conditions
	// u.Where(u.ID.Eq(2)).Update(u.Age, u.Age.Add(1))
	// or
	// u.Where(u.ID.Eq(2)).UpdateSimple(u.Age.Add(1))

	u.Where(u.ID.Eq(2)).UpdateSimple(u.Age.Zero())
}

func updateMultiFields() {
	u := dao.User

	newU := map[string]any{
		"Age": 8,
	}
	u.Where(u.ID.Eq(2)).Updates(newU)
}

func deleteOne() {
	u := dao.User

	u.Where(u.ID.Eq(13)).Delete()
}

func selectUsers() {
	u := dao.User
	users, err := u.Where(u.ID.In(1, 2, 3)).Find()
	if err != nil {
		log.Println(err)
	}

	log.Println(users)
}

func selectUserPage() {
	u := dao.User

	users, count, err := u.Where(u.ID.Gt(2)).Order(u.Age.Desc()).FindByPage(2, 2)
	if err != nil {
		log.Println(err)
	}

	log.Println(count)
	log.Println(users)
}

func transaction() {
	err := dao.Q.Transaction(func(tx *dao.Query) error {
		res, err := tx.User.Where(tx.User.ID.Eq(2)).UpdateSimple(tx.User.Age.Add(1))
		if err != nil {
			return err
		}
		log.Println(res)
		res, err = tx.User.Where(tx.User.ID.Eq(2)).UpdateSimple(tx.User.ID.Add(1))
		if err != nil {
			return err
		}
		log.Println(res)
		return nil
	})

	if err != nil {
		log.Println(err)
	}
}

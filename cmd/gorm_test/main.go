package main

import (
	gormdb "gowebdemo/internal/pkg/gorm_db"
	"log"

	"gorm.io/gorm"
)

type Company struct {
	ID   int
	Name string
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

type Luanguage struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_languages;"`
}

type User struct {
	gorm.Model
	Name        string
	CompanyID   int
	Company     Company
	CreditCards []CreditCard
	Languages   []Luanguage `gorm:"many2many:user_languages;"`
}

func main() {
	db, err := gormdb.Connect()
	if err != nil {
		panic("connect db err")
	}

	db.AutoMigrate(&User{}, &Company{}, &CreditCard{})

	// company := Company{
	// 	Name: "NB",
	// }

	// var company Company
	// db.First(&company, "ID", "1")

	// lan := Luanguage{
	// 	Name: "chinese",
	// }

	// db.Create(&lan)

	// card := CreditCard{
	// 	Number: "6550001",
	// }

	// db.Create(&card)

	// var u User
	// err = db.Model(&User{}).Preload("CreditCards").Preload("Company").Where("ID = 1").First(&u).Error
	// if err != nil {
	// 	log.Println(err)
	// }

	// log.Printf("user: %v", u)

	// var lan Luanguage
	// db.Where("id = 1").First(&lan)

	// u := &User{
	// 	Name:        "u8",
	// 	CompanyID:   1,
	// 	CreditCards: []CreditCard{card},
	// 	Languages:   []Luanguage{lan},
	// }

	// res := db.Create(u)
	// if res.Error != nil {
	// 	log.Println(res.Error)
	// }

	// log.Println(res.RowsAffected)

	u := User{}
	db.Where("id = 6").First(&u)
	log.Println(u)
}

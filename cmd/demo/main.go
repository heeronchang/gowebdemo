package main

import (
	gormdb "gowebdemo/internal/pkg/gorm_db"
	"time"

	"gorm.io/gorm"
)

// func main() {
// 	var wg sync.WaitGroup
// 	var str atomic.Value // 定义一个原子变量
// 	str.Store("hello, world")
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			oldStr := str.Load().(string) // 读取原子变量的值
// 			if i%2 == 0 {
// 				newStr := oldStr + "!"
// 				str.Store(newStr) // 写入原子变量的值
// 				fmt.Println(newStr)
// 			} else {
// 				fmt.Println(oldStr)
// 			}
// 		}(i)
// 	}
// 	wg.Wait()
// }

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	gorm.Model
	Name      string `gorm:"column:user_name"`
	Age       int8
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt int64      `gorm:"autoUpdateTime:nano"`
	Updated   int64      `gorm:"autoUpdateTime"`
	Created   int64      `gorm:"autoCreateTime"`
}

func main() {
	db, err := gormdb.Connect()
	if err != nil {
		panic("connect db err")
	}

	db.AutoMigrate(&Product{}, &User{})
	// db.Create(&Product{Code: "D42", Price: 200})

	// var product Product
	// db.First(&product, "code=?", "D42")

	// db.Model(&product).Update("Price", 100)
	// db.Model(&product).Updates(Product{Price: 400, Code: "F42"})
	// db.Model(&product).Updates(map[string]any{"Price": 100, "Code": "F42"})

	// db.Delete(&product, 1)

	// log.Println(product)

	// var products []Product

	// db.Unscoped().Where("deleted_at is not null").Find(&products)

	// log.Println(products)

	db.Create(&User{Name: "Tom", Age: 18})
	// var u = &User{
	// 	Name: "Jack",
	// 	Age:  17,
	// }
	// u.UpdatedAt = time.Now()

	// db.Create(u)
}

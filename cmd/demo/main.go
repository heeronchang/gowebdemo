package main

import (
	"errors"
	"fmt"
	gormdb "gowebdemo/internal/pkg/gorm_db"
	"time"

	"github.com/google/uuid"
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
	Name       string `gorm:"column:user_name"`
	Age        int8
	UUID       string
	Role       int8
	CreditCard CreditCard
	CreatedAt  *time.Time `gorm:"autoCreateTime"`
	UpdatedAt  int64      `gorm:"autoUpdateTime:nano"`
	Updated    int64      `gorm:"autoUpdateTime"`
	Created    int64      `gorm:"autoCreateTime"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New().String()
	if u.Role == 1 {
		return errors.New("user is admin")
	}
	fmt.Println(u)
	return
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

func main() {
	db, err := gormdb.Connect()
	if err != nil {
		panic("connect db err")
	}

	db.AutoMigrate(&Product{}, &User{}, &CreditCard{})
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

	// db.Create(&User{Name: "Tom", Age: 18})
	// var u = &User{
	// 	Name: "Jack",
	// 	Age:  17,
	// }
	// u.UpdatedAt = time.Now()

	// db.Create(u)

	// // 创建多条记录
	// users := []*User{
	// 	{Name: "Jinzhu", Age: 18},
	// 	{Name: "Jackson", Age: 19},
	// }

	// result := db.Create(users) // pass a slice to insert multiple row
	// fmt.Println(result)

	// // 插入指定字段的记录
	// user := &User{
	// 	Name: "Tom",
	// 	Age:  10,
	// }

	// db.Select("Name", "CreateAt").Create(user)

	// // 插入忽略指定字段的记录
	// user := &User{
	// 	Name: "Tomm",
	// 	Age:  8,
	// }

	// db.Omit("Age").Create(user)

	// // 创建勾子
	// // GORM allows user defined hooks to be implemented for
	// // BeforeSave, BeforeCreate, AfterSave, AfterCreate.
	// // These hook method will be called when creating a record,
	// // refer Hooks for details on the lifecycle
	// user := &User{
	// 	Name: "Jackson",
	// 	Age:  20,
	// 	Role: 2,
	// }

	// result := db.Create(user)
	// fmt.Println(result.Error)

	// // 忽略勾子
	// // If you want to skip Hooks methods, you can use the SkipHooks session mode
	// user := &User{
	// 	Name: "Jack",
	// 	Age:  20,
	// 	Role: 1,
	// }
	// db.Session(&gorm.Session{SkipHooks: true}).Create(user)

	// // 根据map创建
	// NOTE When creating from map, hooks won’t be invoked,
	// associations won’t be saved and primary key values won’t be back filled
	// db.Model(&User{}).Create(map[string]any{"Name": "Lily", "Age": 17, "Role": 2})

	// 使用 SQL 表达式 创建记录
	// db.Model(User{}).Create(map[string]interface{}{
	// 	"Name":     "jinzhu",
	// 	"Location": clause.Expr{SQL: "ST_PointFromText(?)", Vars: []interface{}{"POINT(100 100)"}},
	// })

	// // 关联创建
	// user := &User{
	// 	Name:       "Heeron",
	// 	Age:        33,
	// 	Role:       2,
	// 	CreditCard: CreditCard{Number: "411111111111"},
	// }
	// db.Create(user)

	// // // skip saving associations with Select, Omit
	// // db.Omit("CreditCard").Create(&user)

	// // // skip all associations
	// // db.Omit(clause.Associations).Create(&user)

	// var user = User{}
	// user.ID = 1
	// db.First(&user)
}

package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func main()  {


	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:123456@tcp(42.193.97.239:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("init db error: %s", err)
	}

	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	result := db.Table("user").Create(&user) // 通过数据的指针来创建
	fmt.Println(user.ID)
	fmt.Println(user.Name)
	fmt.Println(result.RowsAffected)
}

type User struct {
	ID int `gorm:"primary_key"`
	Name string
	Age int
	Birthday time.Time
}

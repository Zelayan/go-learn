package main

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	gorm.Model
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:111111@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	panic(err)
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
	db.Transaction(func(tx *gorm.DB) error {
		tx.Model(&User{}).Where("name = ?", "chenze").Update("age", 11)
		//panic("err")
		return nil
	})
}
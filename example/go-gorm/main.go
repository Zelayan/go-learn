package main

import (
	"errors"
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       uint
	Name     string
	Email    *string
	Age      uint8
	Birthday *time.Time
	gorm.Model
}

var db = initDb()

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

	//autoTransaction()
	//manualTransaction()
	nestedTransaction()
}

func nestedTransaction() error {
	tx := db.Begin()
	tx.Transaction(func(tx *gorm.DB) error {
		tx.Model(&User{}).Where("name = ?", "chenze").Update("age", "1")
		tx.Transaction(func(tx *gorm.DB) error {
			tx.Model(&User{}).Where("name = ?", "chenze").Update("age", "2")
			return errors.New("rollback 2")
		})
		tx.Transaction(func(tx *gorm.DB) error {
			tx.Model(&User{}).Where("name = ?", "chenze").Update("age", "3")
			return nil
		})
		return nil
	})
	return nil
}

func manualTransaction() error {
	// 开始事物
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	tx.Model(&User{}).Where("name = ?", "chenzeze").Update("age", "10")
	affected := tx.Model(&User{}).Where("name = ?", "chenzeze").Update("name", "chenze01").RowsAffected
	if affected == 0 {
		tx.Rollback()
		return fmt.Errorf("update name faield")
	}
	tx.Commit()
	return nil
}

// 自动事物
func autoTransaction() error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recory")
		}
	}()
	err := db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&User{
			Name:     "1",
			Email:    nil,
			Age:      0,
			Birthday: nil,
			Model:    gorm.Model{},
		})
		panic("xxxxx")
		return nil
	})
	return err
}

func initDb() *gorm.DB {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)

	dsn := "root:111111@tcp(zeze.com:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	/*err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}*/
	return db
}

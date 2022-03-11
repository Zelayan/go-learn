package main

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID int `gorm:"primary_key"`
	Name string
	Age int
	Birthday time.Time
}

type DB struct {
	db *gorm.DB
}

func (d *DB)InitTable() error {
	 return d.db.AutoMigrate(User{})
}

func (d *DB) Create(user *User) error{
	return d.db.Create(user).Error
}
package main

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestInitTable(t *testing.T) {
	check := assert.New(t)
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	check.Nil(err)

	d := &DB{db: db}
	err = d.InitTable()
	check.Nil(err)
	err = d.Create(&User{
		ID:       1,
		Name:     "xx",
		Age:      0,
		Birthday: time.Time{},
	})
	check.Nil(err)
}
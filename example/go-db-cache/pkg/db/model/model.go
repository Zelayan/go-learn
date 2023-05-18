package model

import "time"

type User struct {
	Model
	Name   string `gorm:"index:idx_name,unique" json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

type Model struct {
	Id              int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	GmtCreate       time.Time `json:"gmt_create"`
	GmtModified     time.Time `json:"gmt_modified"`
	ResourceVersion int64     `json:"resource_version"`
}

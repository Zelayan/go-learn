package db

import (
	"go-db-cache/pkg/db/user"
	"gorm.io/gorm"
)

type ShareDaoFactory interface {
	User() user.UserInterface
}

type shareDaoFactory struct {
	db *gorm.DB
}

func (s *shareDaoFactory) User() user.UserInterface {
	return user.NewUser(s.db)
}

func NewShareDaoFactory(db *gorm.DB) ShareDaoFactory {
	return &shareDaoFactory{db: db}
}

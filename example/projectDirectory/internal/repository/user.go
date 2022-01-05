package repository

import (
	"project-directory/internal/datastruct"
	"project-directory/internal/dto"
)

type UserQuery interface {
	CreateUser(user datastruct.Person) (*int64, error)
	GetUser(id int64) (*datastruct.Person, error)
	DeleteUser(userID int64) error
	UpdateUser(person dto.Person) (*datastruct.Person, error)
	GetUserPasswordByEmail(email string) (*string, error)
	GetEmailByUserID(id int64) (string, error)
	GetEmailCode(id int64) (*int64, error)
	UpdateEmailCode(userID, code int64) error
	VerifiedTrueEmailCodeZero(id int64) error
	GetUserIdByEmail(email string) (*int64, error)
}

package user

import (
	"context"
	"go-db-cache/pkg/db/model"
	"gorm.io/gorm"
	"time"
)

type UserInterface interface {
	Create(ctx context.Context, obj *model.User) (*model.User, error)
	Update(ctx context.Context, uid int64, resourceVersion int64, updates map[string]interface{}) error
	Delete(ctx context.Context, uid int64) error
	Get(ctx context.Context, uid int64) (*model.User, error)

	GetByName(ctx context.Context, name string) (*model.User, error)
}

type user struct {
	db *gorm.DB
}

// Create 创建用户
func (u *user) Create(ctx context.Context, obj *model.User) (*model.User, error) {
	now := time.Now()
	obj.GmtCreate = now
	obj.GmtModified = now
	if err := u.db.Create(obj).Error; err != nil {
		return nil, err
	}
	return obj, nil
}

func (u user) Update(ctx context.Context, uid int64, resourceVersion int64, updates map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (u user) Delete(ctx context.Context, uid int64) error {
	//TODO implement me
	panic("implement me")
}

func (u user) Get(ctx context.Context, uid int64) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u user) GetByName(ctx context.Context, name string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUser(db *gorm.DB) UserInterface {
	return &user{db: db}
}

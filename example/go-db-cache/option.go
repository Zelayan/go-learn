package main

import (
	"go-db-cache/pkg/db"
	"go-db-cache/pkg/db/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Options struct {
	DB      *gorm.DB
	Factory db.ShareDaoFactory
}

func NewOptioon() (*Options, error) {
	return &Options{}, nil
}

func (o *Options) Complete() error {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	dsn, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	dsn.AutoMigrate(model.User{})
	o.Factory = db.NewShareDaoFactory(dsn)
	return err
}

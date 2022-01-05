package repository

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type DAO interface {
	NewUserQuery() UserQuery
}

type dao struct{}


var DB *sql.DB

func NewDB() (*sql.DB, error) {
	viper.AddConfigPath("../config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("cannot read from a config")
	}
	host := viper.Get("database.host").(string)
	port := viper.Get("database.port").(int)
	user := viper.Get("database.user").(string)
	dbname := viper.Get("database.dbname").(string)
	password := viper.Get("database.password").(string)

	// Starting a database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return DB, nil
}

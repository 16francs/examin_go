package datastore

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/16francs/examin_go/config"
)

// Connect - DBへの接続
func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", dbConfig())
	if err != nil {
		panic("Failed to connect database")
	}

	db.LogMode(true)
	return db
}

// dbConfig - DBの設定
func dbConfig() string {
	env, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("alert: %s", err)
	}

	username := env.DBUser
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	database := env.DBName

	connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)

	return connect
}

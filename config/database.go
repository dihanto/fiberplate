package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "dihanto:jangkrik@tcp(127.0.0.1:3306)/fiberplate?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println("success connect database")
	return connect(dsn)
}

func connect(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		log.Println("[DATABASE]:: CONNECTION_ERROR")
		panic(err)
	}

	return db
}

func Disconnect(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("[DATABASE]::DISCONNECTION_ERROR")
		panic(err)
	}
	sqlDB.Close()
}

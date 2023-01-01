package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local" // it is for test!
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if sqlDB, err := db.DB(); err != nil {
		panic(err)
	} else {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	db.AutoMigrate(&User{})

}

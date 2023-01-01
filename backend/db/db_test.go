package db

import (
	"fmt"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt time.Time
	Code      string
	Price     uint
	List      []string
}

func TestMain(t *testing.T) {
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "D42", Price: 100, List: []string{"a", "b", "c"}})

	var product Product
	db.First(&product, 1) // find product with integer primary key
	fmt.Println(product)
	db.First(&product, "code = ?", "D42") // find product with code D42
	fmt.Println(product)

	// db.Model(&product).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// db.Delete(&product, 1)
}

package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	// dsn := "fengfeng:root@tcp(127.0.0.1:5432)/db?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "host=127.0.0.1 user=fengfeng password=root dbname=db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}

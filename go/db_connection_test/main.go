package main

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func main() {
	dsn := "dev:dev@tcp(localhost:53306)/ccrs_db?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	println("connected to database")
}

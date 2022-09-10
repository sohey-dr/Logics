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

	rows, err := db.Raw("show tables").Rows()
	if err != nil {
			panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
			var table string
			if err := rows.Scan(&table); err != nil {
					panic(err.Error())
			}
			fmt.Println(table)
	}
}

package main

import (
	"fmt"
	"gorm.io/gorm"
)
import "gorm.io/driver/mysql"

func main() {
	dsn := "root:003416nba@tcp(127.0.0.1:3306)/orm_test?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("conn suc ...")
}

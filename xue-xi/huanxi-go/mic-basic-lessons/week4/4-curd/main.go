package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"mic-basic-lessons/week4/3-zidongqianyi/model"
	"os"
	"time"
)

var db *gorm.DB

func init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	var err error
	dsn := "root:003416nba@tcp(127.0.0.1:3306)/orm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("conn suc ...")
}

func main() {
	//db.Create(&model.Product{Code: "D42", Price: 200})

	var p model.Product
	//db.First(&p , 1)
	//db.First(&p , "code=?" , "D42")
	db.First(&p, 1)
	db.Model(&p).Updates(model.Product{
		Code:  "FF42",
		Price: 600,
	})
	db.First(&p, 1)
	fmt.Println(p.Code, p.Price)

}

package main

import (
	"fmt"
	"go-2022/xue-xi/huanxi-go/mic-basic-lessons/week4/12-gorm-guanlian/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
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
	//db.AutoMigrate(&model.Company{})
	//db.AutoMigrate(&model.Employer{})
}

func add(){
	c1 := model.Company{
		Name: "面向加薪1",
	}
	/*e1 := model.Employer{
		Name:    "欢喜啊",
		Company: c1,
	}*/
	e1 := model.Employer{
		Name:    "22",
		Company: c1,
	}
	db.Create(&e1)
}

func cx1() {
	var e1 model.Employer
	/*db.First(&e1)
	fmt.Println(e1.Name)
	fmt.Println(e1.Company.ID)
	fmt.Println(e1.Company.Name)*/

	/*db.Preload("Company").First(&e1)
	fmt.Println(e1.Name)
	fmt.Println(e1.Company.ID)
	fmt.Println(e1.Company.Name)*/

	db.Joins("Company").First(&e1)
	fmt.Println(e1.Name)
	fmt.Println(e1.Company.ID)
	fmt.Println(e1.Company.Name)
}

func main() {
	add()
	//cx1()
}

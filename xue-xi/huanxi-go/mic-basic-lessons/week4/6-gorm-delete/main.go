package main

import (
	"fmt"
	"go-2022/xue-xi/huanxi-go/mic-basic-lessons/week4/6-gorm-delete/model"
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
	//db.AutoMigrate(&model.Food{})
}

func inset() {
	f1 := model.Food{
		Name: "茄子",
		Tag:  1,
	}
	f2 := model.Food{
		Name: "土豆丝",
		Tag:  2,
	}
	f3 := model.Food{
		Name: "萝卜",
		Tag:  3,
	}
	var foodList []model.Food
	{
	}

	foodList = append(foodList, f1)
	foodList = append(foodList, f2)
	foodList = append(foodList, f3)

	db.CreateInBatches(foodList, 3)
}

func main() {
	//inset()
	//db.Model()
	var p model.Food
	db.First(&p, 2)
	//fmt.Println(p.Tag,p.ID,p.Name)
	db.Delete(&p)

	//批量删除
	db.Where("tag=?" , 777).Delete(&model.Food{})
}

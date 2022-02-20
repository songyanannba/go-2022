package internal

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"mic-training-lessons-part2/model"
	"os"
	"time"
)

var DB *gorm.DB
var err error

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"dbName"`
	UserName string `mapstructure:"userName"`
	Password string `mapstructure:"password"`
}

func InitDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Warn, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	var err error
	//dsn := "root:003416nba@tcp(127.0.0.1:3306)/happy_account_mic_traning?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		AppConf.DBConfig.UserName,
		AppConf.DBConfig.Password,
		AppConf.DBConfig.Host,
		AppConf.DBConfig.Port,
		AppConf.DBConfig.DBName,
	)
	zap.S().Infof(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}
	fmt.Println("conn suc ...")
	//err = DB.AutoMigrate(&model.Product{},&model.Category{},&model.Brand{},&model.Advertise{})
	err = DB.AutoMigrate(&model.Category{}, &model.Brand{}, &model.Advertise{}, &model.Product{})
	if err != nil {
		fmt.Println(err)
	}
}

func MyPaging(pageNo, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNo < 1 {
			pageNo = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize < 1:
		}
		offset := (pageNo - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

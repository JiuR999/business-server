package db

import (
	"BusinessServer/env"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type orm struct {
	Engine *gorm.DB
}

var Orm = new(orm)

func init() {
	userName := env.GetConfig().UserName
	password := env.GetConfig().Password
	database := env.GetConfig().Name
	address := env.GetConfig().Address
	port := env.GetConfig().DataBasePort
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, address, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 设置日志级别为Info，以便打印出所有SQL语句
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // 日志输出的目标，前缀和日志包含的内容
			logger.Config{
				//SlowThreshold:             500 * time.Millisecond,           // 慢SQL阈值
				LogLevel:                  logger.Info, // 日志级别
				IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  true,        // 使用彩色打印
			},
		),
	})

	if err == nil {
		fmt.Println("链接成功")
		Orm.Engine = db
	}

}

func (d *orm) DB() *gorm.DB {
	return d.Engine
}

package model

import (
	"fmt"
	"go_bin_blog/utils"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var err error

func InitDb() {

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName)
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		//gorm日志模式:silent
		Logger: logger.Default.LogMode(logger.Silent),
		//外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		//禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			//使用单数表名，启用该选项，此时`User` 的表名是`user`
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数:", err)
		os.Exit(1)
	}

	//禁用默认表名的复数
	//db.SingularTable(true)

	_ = db.AutoMigrate(&User{}, &Article{}, &Category{})

	sqlDB, _:= db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	//sqlDB.Close()
}

package db

import (
	"fmt"
	"heavenMs-NAP-golang/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	GormDB *gorm.DB
)

func NewGormConn() error {
	// 定义数据库连接字符串
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local&timeout=%v",
		config.Cfg.Mysql.UserName,
		config.Cfg.Mysql.Password,
		config.Cfg.Mysql.Host,
		config.Cfg.Mysql.Port,
		config.Cfg.Mysql.Database,
		config.Cfg.Mysql.Conntime)
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	GormDB = db

	// TODO: model
	GormDB.AutoMigrate()

	// TODO: 自动迁移，执行一些sql文件
	return nil
}

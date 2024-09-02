package db

import (
	"embed"
	"errors"
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"heavenMs-NAP-golang/config"
	"heavenMs-NAP-golang/internal/model"
	"heavenMs-NAP-golang/utils"
)

var (
	GormDB *gorm.DB
	//go:embed migration/*.sql
	DBHistory embed.FS
)

func NewGormConn() error {
	logrus.Info("[database] 开始初始化")
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

	// TODO: 建表语句先放sql文件里，后续在放入代码中
	// GormDB.AutoMigrate(
	// 	model.Migration{},
	// 	model.Account{},
	// 	model.Alliance{},
	// 	model.Allianceguilds{},
	// 	model.AreaInfo{},
	// 	model.BbsReplies{},
	// 	model.BbsThreads{},
	// 	model.BosslogDaily{},
	// 	model.BosslogWeekly{},
	// 	model.Buddies{},
	// 	model.Characters{},
	// 	model.Cooldowns{},
	// 	model.DropDataGlobal{},
	// 	model.Dueyitems{},
	// 	model.Dueypackages{},
	// 	model.Eventstats{},
	// 	model.Famelog{},
	// 	model.Skillmacros{},
	// 	model.Skills{},
	// 	model.Specialcashitems{},
	// 	model.Storages{},
	// 	model.TempData{},
	// 	model.Trocklocations{},
	// 	model.Wishlists{},
	// 	model.Worldtransfers{},
	// )
	if err := checkTable(GormDB); err != nil {
		return err
	}

	return nil
}

// 检查数据库中表结构是否是最新
func checkTable(db *gorm.DB) error {
	// 先查表,找已执行记录
	var mgDB []model.Migration
	db.Order("id asc").Find(&mgDB)

	mgFile, err := getFileList()
	if err != nil {
		return err
	}
	if len(mgDB) > len(mgFile) {
		return errors.New("the file version is low in db, please check the file version")
	}
	for i, v := range mgFile {
		if len(mgDB) > i {
			if v.Tag != mgDB[i].Tag || v.Md5 != mgDB[i].Md5 {
				return fmt.Errorf("in the %d_%s.sql,files and db history inconsistent", v.Tag, v.Name)
			}
			continue
		}
		// 执行新 sql
		if err := db.Exec(v.Content).Error; err != nil {
			return err
		}
		// 更新记录至 db
		v.ExecTime = time.Now()
		v.IsSuccess = true
		db.Create(&v)
	}

	return nil
}

func getFileList() ([]model.Migration, error) {
	de, err := DBHistory.ReadDir("migration")
	if err != nil {
		return nil, err
	}
	var mg model.MigrationSlice

	for _, v := range de {
		s := strings.Split(strings.TrimSuffix(v.Name(), ".sql"), "_")
		if len(s) == 0 {
			continue
		}
		i, err := strconv.Atoi(s[0])
		if err != nil || i == 0 {
			continue
		}

		b, err := DBHistory.ReadFile(filepath.Join("migration", v.Name()))
		if err != nil {
			return nil, err
		}
		if s2 := utils.CalcStringMD5(string(b)); s2 != "" {
			if len(s) == 1 {
				mg = append(mg, model.NewMigration(i, "", s2, string(b)))
			} else {
				mg = append(mg, model.NewMigration(i, s[1], s2, string(b)))
			}
		}
	}

	sort.Sort(mg)

	return mg, nil
}

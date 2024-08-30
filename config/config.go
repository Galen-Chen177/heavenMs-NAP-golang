package config

import (
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ReadConfig(cfgFile string) {
	logrus.Info("[config] 开始初始化")
	if !strings.HasSuffix(cfgFile, ".yaml") {
		logrus.Fatalf("仅支持yaml文件")
	}
	dir, file := filepath.Split(cfgFile)
	file = file[:len(file)-5]

	viper.SetConfigName(file)   // 配置文件名 (不包括扩展名)
	viper.SetConfigType("yaml") // 如果配置文件的扩展名不是 yaml ，需要指定
	viper.AddConfigPath(dir)
	viper.AddConfigPath(".") // 也可以添加多个路径

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Error reading config file, %s", err)
	}

	// 将配置映射到结构体
	if err := viper.Unmarshal(&Cfg); err != nil {
		logrus.Fatalf("Unable to decode into struct, %v", err)
	}

	// 监控配置文件变化并热更新
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Infof("Config file changed: %s", e.Name)
		if err := viper.Unmarshal(&Cfg); err != nil {
			logrus.Errorf("Unable to decode into struct, %v", err)
			return
		}
		logrus.Info("config flush:", Cfg)
	})
}

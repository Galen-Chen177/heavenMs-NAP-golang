package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"heavenMs-NAP-golang/config"
	"heavenMs-NAP-golang/internal/db"
	"heavenMs-NAP-golang/internal/manager"
	"heavenMs-NAP-golang/internal/wz"
)

func main() {
	app := &cli.App{
		Name:  "heavenMs-NAP-golang",
		Usage: "heavenMs-NAP-golang usgae",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Usage:   "listen port",
				Aliases: []string{"p"},
			},
			&cli.BoolFlag{
				Name:  "http",
				Usage: "change to http server model",
			},
			&cli.StringFlag{
				Name:     "config",
				Usage:    "config file path",
				Aliases:  []string{"c"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "wzfile",
				Usage:    "wz file dir",
				Aliases:  []string{"wz"},
				Required: true,
			},
		},
		Action: AppAction,
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Error("app Run err:", err)
	}
}

func AppAction(ctx *cli.Context) error {
	if err := mainInit(ctx.String("config"), ctx.String("wzfile")); err != nil {
		logrus.Fatalf("main init err:%v", err)
	}

	if err := manager.NewManager().Run(); err != nil {
		logrus.Error("manager run err:", err)
		return err
	}

	return nil
}

func mainInit(cfgFile, wzFile string) error {
	start := time.Now()
	defer func() {
		logrus.Info("数据加载耗时:", time.Since(start))
	}()
	// 加载配置
	config.ReadConfig(cfgFile)
	// 初始化数据库
	if err := db.NewGormConn(); err != nil {
		return err
	}
	// 加载wz文件
	wz.InitWzManage(wzFile)

	return nil
}

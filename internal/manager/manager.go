package manager

import (
	"github.com/sirupsen/logrus"

	"heavenMs-NAP-golang/config"
	"heavenMs-NAP-golang/internal/service"
)

type IManager interface {
	Run() error
}

type Manager struct {
	dbSer      service.IDbService
	loginSer   service.ILoginService
	mallSer    service.IMallService
	channelSer service.IChannelService
}

func NewManager() IManager {
	return &Manager{
		dbSer:      service.NewDbService(),
		loginSer:   service.NewLoginService(),
		mallSer:    service.NewMallService(),
		channelSer: service.NewChannelService(),
	}
}

func (m *Manager) Run() error {
	if config.Cfg.App.Admin {
		logrus.Info("[登录模式] 只允许管理员登录")
	} else {
		logrus.Info("[登录模式] 允许所有人登录")
	}

	if config.Cfg.Server.Register.Auto {
		logrus.Info("[注册模式] 自动生成账号")
	} else {
		logrus.Info("[注册模式] 手动生成账户,包括[网页注册/GM工具注册]")
	}

	logrus.Info("[账户信息] 重置所有玩家状态为未登录")
	_, err := m.dbSer.ResetAllLoginStatus()
	if err != nil {
		logrus.Fatal("重置状态失败:", err)
	}
	_, err = m.dbSer.ResetAllLastGainHiredMerchant()
	if err != nil {
		logrus.Fatal("重置状态失败:", err)
	}
	// TODO
	// 清空过期点券
	// 载入点券比例
	// 载入未失效点券
	// 加载最大的现金id
	// 添加定时任务
	logrus.Info("[登录服务器] 开始初始化")

	return nil
}

package config

var Cfg Config

type Mysql struct {
	UserName string
	Password string
	Host     string
	Port     int
	Database string
	Conntime string
}

type Config struct {
	App    App
	Mysql  Mysql
	Server Server
}

type App struct {
	Port     int
	UserName string
	PassWord string
	SerAddr  string

	// Debug 调试配置
	Debug struct {
		// Enabled 是否启用调试
		Enabled bool
	}

	// Logger 日志配置
	Logger struct {
		// Packet 封包配置
		Packet struct {
			// Display 显示封包日志
			Display bool

			// Debug 显示调试封包日志
			Debug bool
		}
	}
	// Admin 是否仅允许管理员登录
	Admin bool
}

type Server struct {
	// Name 服务器名称
	Name string

	// Flag 游戏标致
	Flag int

	// Limit 游戏限制
	Limit struct {
		// Online 在线人数
		Online int

		// Characters 角色数量
		Characters int
	}

	// Register 注册设置
	Register struct {
		// Auto 开启自动注册
		Auto bool
	}

	// Rand 随机
	Rand struct {
		// Drop 是否随机掉落物品
		Drop bool
	}

	// Login 登陆服务设置
	Login struct {
		// Message 上线后发给玩家的消息
		Message struct {
			// Content 消息内容
			Content string

			// Event 消息事件
			Event string
		}
	}

	// Channel 频道设置
	Channel struct {
		// Port 频道起始端口
		Port int

		// Count 频道数量，遍历生成以起始端口为基础的多个频道端口
		Count int
	}

	// World 世界设置
	World struct {
		// Rate 倍数
		Rate struct {
			// Exp 经验倍数
			Exp int

			// Gold 金币倍数
			Gold int

			// Drop 物品掉落倍数
			Drop struct {
				// Normal 普通物品掉落倍数
				Normal int

				// Boss 物品掉落倍数
				Boss int
			}

			// Cash 点卷掉落倍数
			Cash int
		}

		// Flags (WFlags) 世界标志,??我也不知道是什么
		Flags int
	}

	// Mall 商城服务配置
	Mall struct {
		// Port 服务端口
		Port int

		// Disabled  商城物品屏蔽ID
		Disabled []string
	}

	// Job 职业群配置
	Job struct {
		// Adventurer 冒险家职业群开关
		Adventurer bool

		// Knights 骑士职业群开关
		Knights bool

		// WarGod  战神职业群开关
		WarGod bool
	}

	// Events 启用的事件列表
	Events []string
}

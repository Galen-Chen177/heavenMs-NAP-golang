package config

var Cfg Config

type Config struct {
	App   App
	Mysql Mysql
}

type App struct {
	Port     int
	UserName string
	PassWord string
	SerAddr  string
}

type Mysql struct {
	UserName string
	Password string
	Host     string
	Port     int
	Database string
	Conntime string
}

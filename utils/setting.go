package utils

import (
	"gopkg.in/ini.v1"
	_ "gopkg.in/ini.v1"
)

var Conf = new(Config)

type Config struct {
	AppName  string `ini:"app_name"`
	Mode     string `ini:"mode"`
	HTTPPort string `ini:"http_port"`

	MySQL MySQLConfig `ini:"mysql"`
	Qiniu QiniuConfig `ini:"qiniu"`
}

type QiniuConfig struct {
	AccessKey   string `ini:"AccessKey"`
	SecretKey   string `ini:"SecretKey"`
	Bucket      string `ini:"Bucket"`
	QiniuServer string `ini:"QiniuServer"`
}

type MySQLConfig struct {
	IP       string `ini:"ip"`
	Port     string `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

func Init(file string) error {
	return ini.MapTo(Conf, file)
}

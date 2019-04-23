package conf

import (
	"github.com/go-ini/ini"
	"log"
)

var cfg *ini.File

type appConfig struct {
	Page   int    `ini:"page"`
	Secret string `ini:"secret"`
}

type mysqlConfig struct {
	User      string `ini:"user"`
	Password  string `ini:"password"`
	Db        string `ini:"db"`
	Host      string `ini:"host"`
	Charset   string `ini:"charset"`
	ParseTime bool   `ini:"parseTime"`
	Loc       string `ini:"loc"`
}

var AppCfg appConfig

var MysqlCfg mysqlConfig

//初始化配置
func InitConfig() {
	cfg, err := ini.Load("conf/conf.ini")

	if err != nil {
		log.Println(err)
	}
	err = cfg.Section("app").MapTo(&AppCfg)
	if err != nil {
		log.Println(err)
	}

	err = cfg.Section("mysql").MapTo(&MysqlCfg)
	if err != nil {
		log.Println(err)
	}
}

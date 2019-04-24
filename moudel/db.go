package moudel

import (
	"api/conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"os"
)

var db *xorm.Engine

func InitDb() {
	var err error
	db, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s",
		conf.MysqlCfg.User,
		conf.MysqlCfg.Password,
		conf.MysqlCfg.Host,
		conf.MysqlCfg.Db,
		conf.MysqlCfg.Charset,
	))

	if err != nil {
		log.Println(err)
		panic(err)
	}
	f, err := os.Create("log/sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	db.SetLogger(xorm.NewSimpleLogger(f))

}

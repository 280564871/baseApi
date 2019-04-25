package models

import (
	"api/conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"log"
	"os"
)

//生成orm 映射struct
//xorm reverse mysql root:root@/shop?charset=utf8 /Users/work/code/go/api/models/tmplate/

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
	log.Println(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s",
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

	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "shop_")
	db.SetTableMapper(tbMapper)

	f, err := os.Create("log/sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	db.SetLogger(xorm.NewSimpleLogger(f))

}

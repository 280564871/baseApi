package pkg

import (
	"api/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func InitDb() {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%v&loc=%s",
		conf.MysqlCfg.User,
		conf.MysqlCfg.Password,
		conf.MysqlCfg.Host,
		conf.MysqlCfg.Db,
		conf.MysqlCfg.Charset,
		conf.MysqlCfg.ParseTime,
		conf.MysqlCfg.Loc,
	))

	if err != nil {
		log.Println(err)
		panic(err)
	}

}

func CloseDb() {
	defer db.Close()
}

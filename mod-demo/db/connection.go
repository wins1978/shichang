package db

import (
	"fmt"
	"log"
	"os"

	"g.demo/config"

	"github.com/jinzhu/gorm"
)

// InitDB 初始化DB (parseTime必须设置为true)
func InitDB() *gorm.DB {
	var source = fmt.Sprintf("%s:%s@tcp(%s:3306)/GO_TESTDB?charset=utf8&parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST)
	var err error
	da, err := gorm.Open("mysql", source)
	if err != nil {
		defer da.Close()
		panic(err)
	}

	da.LogMode(true)
	da.SetLogger(log.New(os.Stdout, "\r\n", 0))
	da.DB().SetMaxIdleConns(10)
	da.DB().SetMaxOpenConns(200)
	//表名不加S后缀
	da.SingularTable(true)

	return da
}

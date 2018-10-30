package main

import (
	"fmt"

	"xorm_demo/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	var source string = fmt.Sprintf("", config.DB_USER, config.DB_PASSWORD, config.DB_HOST)
	engine, err = xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
}

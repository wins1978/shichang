package main

import (
	"github.com/go-xorm/xorm"
	"time"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int64
	Name string
	Salt string
	Age int
	Passwd string `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}	
	
func Fmtdemo() {
	fmt.Println("secod")
	fmt.Println("3")
	user := DB_USER
	pwd := DB_PASSWORD
	host := DB_HOST
	var connStr string = user + ":"+pwd+ "@tcp("+host+":3306)/shichang?charset=utf8&parseTime=True&loc=Local"
	
	//fmt.Println(connStr)
	engine, _ := xorm.NewEngine("mysql", connStr)
	
	engine.Sync2(new(User))
}
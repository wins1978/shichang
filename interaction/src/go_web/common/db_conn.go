package common

import (
//DB Setting
"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/mysql"
"go_web/config"
)

var Db *gorm.DB

func InitDB() {
	var err error
	user := config.DB_USER
	pwd := config.DB_PASSWORD
	host := config.DB_HOST
	var connStr string = user + ":"+pwd+ "@tcp("+host+":3306)/shichang?charset=utf8&parseTime=True&loc=Local"
	Db, err = gorm.Open("mysql", connStr)
    if err != nil {
        panic(err)
    }
	Db.Debug()
	//defer Db.Close()
}
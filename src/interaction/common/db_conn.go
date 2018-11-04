package common

import (
	//DB Setting

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func InitDB() {
	var err error
	user := "root"
	pwd := "Cosfcoco00"
	host := "localhost"
	var connStr string = user + ":" + pwd + "@tcp(" + host + ":3306)/shichang?charset=utf8&parseTime=True&loc=Local"
	Db, err = gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	Db.Debug()
	//defer Db.Close()
}

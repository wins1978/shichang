package main

import (
	"fmt"
	"log"
	"os"

	"g.demo/config"
	"g.demo/model"
	"github.com/jinzhu/gorm"
)

func main() {
	var DB *gorm.DB
	var source = fmt.Sprintf("%s:%s@tcp(%s:3306)/GO_TESTDB?charset=utf8&parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST)
	var err error
	DB, err = gorm.Open("mysql", source)
	if err != nil {
		fmt.Println(err)
		defer DB.Close()
	}

	DB.LogMode(true)
	DB.SetLogger(log.New(os.Stdout, "\r\n", 0))

	var user model.User
	DB.Where(&user).First(&user)

	fmt.Println(user.Name)
	

}

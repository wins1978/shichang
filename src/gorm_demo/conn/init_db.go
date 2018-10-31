package conn

import (
	"fmt"
	"gorm_demo/config"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// InitDB 初始化DB
func InitDB() {
	var source = fmt.Sprintf("%s:%s@tcp(%s:3306)/GO_TESTDB?charset=utf8", config.DB_USER, config.DB_PASSWORD, config.DB_HOST)
	var err error
	DB, err = gorm.Open("mysql", source)
	if err != nil {
		fmt.Println(err)
		defer DB.Close()
	}

	DB.LogMode(true)
	DB.SetLogger(log.New(os.Stdout, "\r\n", 0))
}

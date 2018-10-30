package conn

import (
	"fmt"
	"xorm_demo/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	DB *xorm.Engine
)

func InitDB() {
	fmt.Println("初始化数据库连接")

	var source string = fmt.Sprintf("%s:%s@tcp(%s:3306)/GO_TESTDB?charset=utf8", config.DB_USER, config.DB_PASSWORD, config.DB_HOST)
	fmt.Println(source)
	var err error
	DB, err = xorm.NewEngine("mysql", source)
	if err != nil {
		DB.Close()
		fmt.Println("数据库连接失败")
	}
}

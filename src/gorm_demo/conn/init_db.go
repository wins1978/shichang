package conn

import (
	"bytes"
	"fmt"
	"gorm_demo/config"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// InitDB 初始化DB (parseTime必须设置为true)
func InitDB() {
	var source = fmt.Sprintf("%s:%s@tcp(%s:3306)/GO_TESTDB?charset=utf8&parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST)
	var err error
	DB, err = gorm.Open("mysql", source)
	if err != nil {
		fmt.Println(err)
		defer DB.Close()
	}

	DB.LogMode(true)
	DB.SetLogger(log.New(os.Stdout, "\r\n", 0))
}

//GetGID 获取协程ID，用于并发记录
func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

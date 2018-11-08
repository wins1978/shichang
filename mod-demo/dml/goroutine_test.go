package dml_test

import (
	"testing"
	"time"

	"g.demo/dml"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//协程测试

//多并发插入1000条数据，数据库本身需要设置最大连接数

func TestMultiInsert(t *testing.T) {
	for i := 0; i < 3000; i++ {
		go dml.InsertUser(i, 1)
	}
	time.Sleep(time.Duration(20) * time.Second)

}

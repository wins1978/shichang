package dml_test

import (
	"testing"
"time"
	"g.demo/db"
	"g.demo/dml"
)

//事务测试

func TestTran(t *testing.T) {
	

	for i:=0;i<100;i++ {
		da := db.InitDB()
		tx := da.Begin()
		go dml.TxInsertDept(tx)
		
	}

	time.Sleep(time.Duration(30) * time.Second)
}

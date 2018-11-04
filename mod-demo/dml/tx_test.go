package dml_test

import (
	"testing"
	"fmt"

	"g.demo/db"
	"g.demo/dml"
	"time"
)

//事务测试

func TestTran(t *testing.T) {
	

	for i:=0;i<10;i++ {
		da := db.InitDB()
		tx := da.Begin()
		go dml.TxInsertDept(tx)
		go dml.TxInsertUser(tx)
		time.Sleep(time.Duration(2) * time.Second)
		if(tx.Error != nil) {
		fmt.Println("---------rollback")
		tx.Rollback()
		}else{
			fmt.Println("---------commit")
			tx.Commit()
		}
	}
}

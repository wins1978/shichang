package ddl

import (
	"fmt"
	"testing"
	"xorm_demo/conn"
	"xorm_demo/models"
)

func TestMain(m *testing.M) {
	fmt.Println("Start ddl")
	conn.InitDB()
	m.Run()
	fmt.Println("End ddl")
}

func TestCreateTable(t *testing.T) {
	t.Log("创建用户表")
	//conn.InitDB()
	user := new(models.User)
	dept := new(models.Department)
	if conn.DB == nil {
		t.Error("DB 连接未初始化")
	}
	conn.DB.Sync2(user)
	conn.DB.Sync2(dept)
	conn.DB.Close()
}

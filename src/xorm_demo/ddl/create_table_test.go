package ddl_test

import (
	"testing"
	"xorm_demo/conn"
	"xorm_demo/models"
)

func TestMain(m *testing.M) {
	conn.InitDB()
	m.Run()
	conn.DB.Close()
}

func TestCreateTable(t *testing.T) {
	if conn.DB == nil {
		t.Error("DB 连接未初始化")
	}

	//
	return

	user := new(models.User)
	dept := new(models.Department)

	err := conn.DB.Sync2(user)
	if err != nil {
		t.Error(err)
	}

	err2 := conn.DB.Sync2(dept)
	if err2 != nil {
		t.Error(err2)
	}
}

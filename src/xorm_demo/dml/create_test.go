package dml_test

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

// 插入一行数据
func TestInsertOne(t *testing.T) {
	var user = new(models.User)
	user.ID = 0
	user.Address = "深圳福田中心"
	user.Age = 10
	user.Name = "ccla"
	user.Passwd = "pwd2018"
	user.Phone = "18012345670"

	id, err := conn.DB.InsertOne(user)
	if err != nil {
		t.Error(err)
	}
	t.Logf("ID is:%d", id)
}

// 插入多行数据
func TestInsertMany(t *testing.T) {
	var dept = new(models.Department)

}

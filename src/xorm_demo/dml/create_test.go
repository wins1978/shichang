package dml_test

import (
	"testing"
	"xorm_demo/conn"
	"xorm_demo/models"
)

func TestMain(m *testing.M) {
	conn.InitDB()
	m.Run()
}

//插入一行数据
func TestInsertOne(t *testing.T) {
	var user = new(models.User)
	user.Id = 3
	user.Address = "深圳福田中心"
	user.Age = 10
	user.Name = "ccla"
	user.Passwd = "pwd2018"
	user.Phone = "18012345678"

	conn.DB.InsertOne(user)

}

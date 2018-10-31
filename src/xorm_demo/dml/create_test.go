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
	user.Id = 1
	user.Name = "ccla"

	id, err := conn.DB.InsertOne(user)
	if err != nil {
		t.Error(err)
	}
	t.Logf("ID is:%d", id)
}

// 插入多行数据
func TestInsertMany(t *testing.T) {

	depts := []*models.Department{}

	var d1 = new(models.Department)
	d1.Id = 1

	d1.DeptName = "企业IT"

	var d2 = new(models.Department)
	d2.Id = 2
	d2.DeptName = "游戏"

	depts = append(depts, d1)
	depts = append(depts, d2)

	conn.DB.Insert(depts)
}

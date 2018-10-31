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
	user.Id = 0
	user.Name = "ccla1"

	id, err := conn.DB.InsertOne(user)
	if err != nil {
		t.Error(err)
	}
	t.Logf("ID is:%d", id)

	dept := new(models.Department)
	dept.Id = 0
	dept.DeptName = "dep1"
	dept.Phone = "1000"

	conn.DB.InsertOne(dept)
}

// 插入多行数据(ID自增)
func TestInsertMany(t *testing.T) {
	//插入多条记录并且不使用批量插入，此时实际生成多条插入语句，每条记录均会自动赋予Id值
	users := make([]*models.Department, 2)
	u1 := new(models.Department)
	u1.DeptName = "bbbb3"
	users[0] = u1

	users[1] = new(models.Department)
	users[1].DeptName = "bbbb4"

	conn.DB.Insert(users)

}

// 插入

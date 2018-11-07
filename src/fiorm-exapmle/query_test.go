package exapmle_test

import (
	"fiorm"
	"fiorm/model"
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fiorm.InitDB("mysql", "GO_TESTDB", "localhost", "root", "Cosfcoco00", 3306)
	m.Run()
}

func TestGetItemWhereFirst(t *testing.T) {
	var user model.User

	var qstr = "name = ? or email = ?"
	query := fiorm.Where(qstr, "wins1", "111").
		Select("name, address, id").
		OrderBy("name desc")

	fiorm.GetItemWhereFirst(&user, query)

	fmt.Println(user.ID)
}

func TestGetItemWhere(t *testing.T) {
	var users []model.User

	var qstr = "name = ? or email = ?"
	query := fiorm.Where(qstr, "wins1", "111").
		Select("name, address, id").
		OrderBy("name desc").
		Limit(10)

	var cnt int64
	fiorm.GetItemWhere(&users, query).Count(&cnt)

	fmt.Println(users)
	fmt.Println("--cnt:", cnt)
}

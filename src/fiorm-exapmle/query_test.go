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
	query := fiorm.Where(qstr, "wins1", "111")

	fiorm.DB.GetItemWhereFirst(&user, query)

	fmt.Println(user.ID)
}

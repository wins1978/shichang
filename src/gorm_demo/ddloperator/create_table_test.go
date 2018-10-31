package ddloperator_test

import (
	"gorm_demo/conn"
	"gorm_demo/model"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestMain(m *testing.M) {
	conn.InitDB()
	m.Run()
	defer conn.DB.Close()
}

func TestCreateTable(t *testing.T) {
	var user model.User
	conn.DB.CreateTable(&user)
}

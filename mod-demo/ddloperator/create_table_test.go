package ddloperator_test

import (
	"g.demo/db"
	"g.demo/model"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestMain(m *testing.M) {
	db.InitDB()
	m.Run()
	defer db.DB.Close()
}

func TestCreateTable(t *testing.T) {
	var user model.User
	db.DB.CreateTable(&user)

	var dept model.Department
	db.DB.CreateTable(&dept)
}
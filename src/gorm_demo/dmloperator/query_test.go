package dmloperator_test

import (
	"gorm_demo/conn"
	"gorm_demo/model"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestSelectOneModel(t *testing.T) {
	var user model.User
	conn.DB.Find(&user)

	t.Log("----email:" + user.Address)
}

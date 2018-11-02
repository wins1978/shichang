package dmloperator_test

import (
	"gorm_demo/conn"
	"testing"
)

func TestMain(m *testing.M) {
	conn.InitDB()
	m.Run()
	//defer conn.DB.Close()
}

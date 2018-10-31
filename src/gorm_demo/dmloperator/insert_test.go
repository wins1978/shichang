package dmloperator_test

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

func TestInsertOne(t *testing.T) {
	//sNumber := "100101"
	//var nAge sql.NullInt64
	//nAge.Int64 = 10
	//nAge.Valid = true

	//var birth time.Time
	//birth = time.Now()

	var user model.User
	user.Address = "addr"
	//user.Role = "r1"
	//user.Age = nAge
	//user.Birthday = &birth
	//user.Email = "test3@qq.com"
	//user.IgnoreMe = 10
	//user.MemberNumber = &sNumber
	user.Name = "wins"
	user.Num = 102
	user.ID = 0
	db := conn.DB.Create(&user)
	if db.Error != nil {
		t.Error(db.Error)
	}
}

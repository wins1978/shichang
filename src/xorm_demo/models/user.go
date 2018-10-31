package models

import "time"

// User 用户表
type User struct {
	ID       int64  `xorm:"pk autoincr"`
	Name     string `xorm:"varchar(100)"`
	Phone    string `xorm:"varchar(20) notnull unique"`
	Address  string `xorm:"varchar(200)"`
	DeptID   int64
	DeptName string `xorm:"varchar(100)"`
	Age      int
	Passwd   string    `xorm:"varchar(200)"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}

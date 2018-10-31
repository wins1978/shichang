package models

import "time"

// User 用户表
type User struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	Name       string    `xorm:"VARCHAR(100)"`
	Password   string    `xorm:"VARCHAR(100)"`
	CreateTime time.Time `xorm:"DATETIME"`
	DeptId     int64     `xorm:"BIGINT(20)"`
}

func (User) TableName() string {
	return "User"
}

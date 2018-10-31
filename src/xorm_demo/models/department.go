package models

import "time"

// Department 部门表
type Department struct {
	ID       int64     `xorm:"pk autoincr"`
	DeptName string    `xorm:"varchar(100)"`
	Email    string    `xorm:"varchar(100)"`
	Area     string    `xorm:"varchar(200)"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}

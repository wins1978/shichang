package models

import "time"

type Department struct {
	Id       int64
	DeptName string    `xorm:"varchar(100)"`
	Email    string    `xorm:"varchar(100)"`
	Area     string    `xorm:"varchar(200)"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}

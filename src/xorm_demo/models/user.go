package models

import "time"

type User struct {
	Id      int64
	Name    string `xorm:"varchar(100)"`
	Phone   string `xorm:"varchar(20)"`
	Address string `xorm:"varchar(200)"`
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

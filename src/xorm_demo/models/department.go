package models

import "time"

//工具自动生成方法:使用xorm工具
//安装工具:go get github.com/go-xorm/cmd/xorm
//命令行输入:xorm reverse mysql root:密码@/xorm?charset=utf8 /home/tym/golib/src/github.com/go-xorm/cmd/xorm/templates/goxorm/
//xorm reverse mysql root:itcloud@123@tcp(10.123.17.81:3306)/GO_TESTDB?charset=utf8 D:\\go_projs\\src\\shichang\\src\\github.com\\go-xorm\\cmd\\xorm\\templates\\goxorm
// Department 部门表
type Department struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	DeptName   string    `xorm:"VARCHAR(100)"`
	Phone      string    `xorm:"VARCHAR(20)"`
	CreateTime time.Time `xorm:"DATETIME"`
}

func (Department) TableName() string {
	return "Department"
}

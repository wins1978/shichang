package dmloperator_test

import (
	"fmt"
	"gorm_demo/conn"
	"gorm_demo/model"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//选择一行，返回实体类
func TestSelectRow(t *testing.T) {
	var user model.User
	user.Role = "r1"
	conn.DB.Where(&user).Find(&user)

	t.Log("----email:" + user.Address)
}

//查找多行，返回实体数组
func TestSelectRows(t *testing.T) {
	var users []model.User
	conn.DB.Where("address = ?", "addr1").Find(&users)

	fmt.Println("---------------------Rows:")
	for _, user := range users {
		fmt.Println(user.Address)
	}
}

//IN查询条件
func TestQueryInCondition(t *testing.T) {
	var users []model.User
	var qry []string
	qry = append(qry, "wins1")
	qry = append(qry, "wins2")
	qry = append(qry, "wins3")

	conn.DB.Where("name in (?)", qry).Find(&users)

	for _, user := range users {
		fmt.Println(user.Name)
	}
}

//查询Or关系
func TestQueryOrCondition(t *testing.T) {
	var users []model.User

	conn.DB.Where("name = ?", "wins1").Or("name =?", "wins2").Find(&users)

	for _, user := range users {
		fmt.Println(user.Name)
	}
}

//通过单个ID查询
func TestQueryByPrimaryKey(t *testing.T) {
	var user model.User

	conn.DB.First(&user, 3)
	fmt.Println(user)
}

//查询多个ID
func TestQueryByPrimaryKeys(t *testing.T) {
	var users []model.User
	ids := []int64{1, 2, 3}

	conn.DB.Where(ids).Find(&users)

	for _, user := range users {
		fmt.Println(user)
	}

}

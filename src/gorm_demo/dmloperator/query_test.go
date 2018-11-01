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
	conn.DB.Where(&user).Find(&user)

	t.Log(user.Birthday)
}

//查找多行，返回实体数组
func TestSelectRows(t *testing.T) {
	var users []model.User
	conn.DB.Where("address like ?", "深圳%").Find(&users)

	fmt.Println("---------------------Rows:")
	for _, user := range users {
		fmt.Println(user)
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

//子查询
func TestSubQuery(t *testing.T) {
	//SELECT * FROM go_testdb.user u where u.name='wins1' and
	//dept_id in (select id from go_testdb.department t where t.dept_name like '%企业%')
	subQ := conn.DB.Table("department").Select("id").Where("dept_name like ?", "%企业%").QueryExpr()
	fmt.Println(subQ)

	var users []model.User
	conn.DB.Where("name = ? and dept_id in (?)", "wins1", subQ).Find(&users)

	for _, user := range users {
		fmt.Println(user)
	}
}

//查询视图
func TestQueryView(t *testing.T) {
	var users_v []model.UserInfoView

	conn.DB.Where("dept_name like ?", "%企业%").Find(&users_v)

	for _, user := range users_v {
		fmt.Println(user)
	}
}

//复杂组合查询
func TestComplexQuery(t *testing.T) {
	var users []model.User
	//SELECT * FROM `user`  WHERE (name like ? and null_age < ?) OR (id = ?)
	//ORDER BY id desc[%wins% 40 5]
	conn.DB.Where("name like ? and null_age < ?", "%wins%", 40).
		Or("id = ?", 5).
		Order("id desc").
		Find(&users)

	for _, user := range users {
		fmt.Println(user.ID)
	}
}

//Join查询
func TestJoinQuery(t *testing.T) {
	var users []model.User
	//SELECT user.id,user.name,d.dept_name FROM `user`
	//join department d on user.dept_id = d.id WHERE (user.name like ?)[%wins%]
	conn.DB.Select("user.id,user.name,d.dept_name").
		Joins("join department d on user.dept_id = d.id").
		Where("user.name like ?", "%wins%").
		Find(&users)

	for _, user := range users {
		fmt.Println(user.ID, user.Name, user.DeptName)
	}
}

//转换查询结果到Dto类中
func TestScanDto(t *testing.T) {
	var dtos []model.UserDto
	//SELECT user.id,user.name,d.dept_name FROM `user`
	//left join department d on user.dept_id = d.id WHERE (user.name like ?)[%wins%]
	conn.DB.Table("user").Select("user.id,user.name,d.dept_name").
		Joins("left join department d on user.dept_id = d.id").
		Where("user.name like ?", "%wins%").
		Scan(&dtos)

	for _, dto := range dtos {
		fmt.Println(dto)
	}
}

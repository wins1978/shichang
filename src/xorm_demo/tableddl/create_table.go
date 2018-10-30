package tableddl

import "xorm_demo/conn"
import "xorm_demo/models"
import "fmt"

//创建用户表
func create_user_tb() {
	fmt.Println("创建用户表")
	conn.InitDB()
	user := new(models.User)
	dept := new(models.Department)
	if conn.DB == nil {
		fmt.Println("DB 连接未初始化")
	}
	conn.DB.Sync2(user)
	conn.DB.Sync2(dept)
	conn.DB.Close()
}

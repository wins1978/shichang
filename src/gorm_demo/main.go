package main

import (
	"database/sql"
	"gorm_demo/conn"
	"gorm_demo/model"
)

func main() {
	conn.InitDB()

	var users []model.User

	age := sql.NullInt64{Int64: 999, Valid: true}
	addr := "深圳南山--aa334"

	values := map[string]interface{}{"address": addr, "null_age": age}

	//***********************
	// 不要采用下面的方式更新数据，当第一条语句有错误时，将更新所有数据（除非判断错误并抛出异常）
	//***********************
	conn.DB.Where("id > ", 3).Find(&users)
	conn.DB.Model(&users).Select("null_age", "address").UpdateColumn(values)
}

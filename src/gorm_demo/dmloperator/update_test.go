package dmloperator_test

import (
	"database/sql"
	"gorm_demo/conn"
	"gorm_demo/model"
	"testing"
)

//更新所有字段
//(先取出记录，替换掉新值后，更新所有字段)
func TestUpdateAllField(t *testing.T) {
	var user model.User

	conn.DB.Where(1).First(&user)

	user.NullAge = sql.NullInt64{Int64: 213, Valid: true}
	user.Address = "深圳南山--update2"

	conn.DB.Save(&user)

	//第一步：选择所有查询结果
	//SELECT * FROM `user`  WHERE (`user`.`id` = 1) ORDER BY `user`.`id` ASC LIMIT 1[] 1

	//第二步：更新所有字段的值（只有赋值的会改变）
	//UPDATE `user` SET `created_at` = ?, `updated_at` = ?, `name` = ?, `null_age` = ?, `email` = ?, `null_string` = ?, `address` = ?, `birthday` = ?, `dept_id` = ?
	//WHERE `user`.`id` = 1
	//[2018-11-01 12:49:58 +0000 UTC 2018-11-01 22:49:50.5869681 +0800 CST m=+0.015014801 wins1 {213 true}   深圳南山--update2 {0001-01-01 00:00:00 +0000 UTC false} {1 true} 1]
}

//更新
//Update方法在更新之前，会调用BeforeUpdate、AfterUpdate
//用于更新UpdatedAt字段，如果考虑更好的性能或者不更新UpdatedAt字段，
//可以考虑UpdateColumn、UpdateColumns方法
func TestUpdate(t *testing.T) {
	var user model.User

	conn.DB.Where(2).First(&user)

	user.NullAge = sql.NullInt64{Int64: 66, Valid: true}
	user.Address = "深圳南山--up66"

	//只能更新一个字段
	//conn.DB.Model(&user).Update("null_age", "")
	//UPDATE `user` SET `null_age` = ?, `updated_at` = ?  WHERE `user`.`id` = ?[{215 true} 2018-11-01 22:57:06.3283321 +0800 CST m=+0.017216101 2]

	//用实体类更新(等同于更新所有字段)
	//conn.DB.Model(&user).Updates(&user)
	//UPDATE `user` SET `address` = ?, `created_at` = ?, `dept_id` = ?, `id` = ?, `name` = ?, `null_age` = ?, `updated_at` = ?  WHERE `user`.`id` = ?[深圳南山--update5 2018-11-01 12:50:02 +0000 UTC {2 true} 2 wins1 {216 true} 2018-11-01 23:11:27.21639 +0800 CST m=+0.017016201 2] 1

	//更新特定的字段
	//conn.DB.Model(&user).Select("null_age", "address").Updates(&user)
	//UPDATE `user` SET `address` = ?, `null_age` = ?, `updated_at` = ?  WHERE `user`.`id` = ?[深圳南山--update5 {216 true} 2018-11-01 23:09:55.8102695 +0800 CST m=+0.018017201 2] 1

	//忽略特定字段
	//conn.DB.Model(&user).Omit("null_age", "address").Updates(&user)
	//UPDATE `user` SET `created_at` = ?, `dept_id` = ?, `id` = ?, `name` = ?, `updated_at` = ?  WHERE `user`.`id` = ?[2018-11-01 12:50:02 +0000 UTC {2 true} 2 wins1 2018-11-01 23:12:42.3051292 +0800 CST m=+0.015015401 2]

	// Update multiple attributes with `map`, will only update those changed fields
	//db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
	// UPDATE users SET name='hello', age=18, actived=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
}

//不调用BeforeUpdate的更新
func TestUpdateColumn(t *testing.T) {
	var users []model.User

	age := sql.NullInt64{Int64: 11, Valid: true}
	addr := "深圳南山--up00"

	values := map[string]interface{}{"address": addr, "null_age": age}

	//更新特定的字段
	query := conn.DB.Where("id > ", 3).First(&users)
	//即使注释掉异常处理，后续的语句也不会执行
	if query.Error != nil {
		t.Error(query.Error)
	}
	query.Select("null_age", "address").UpdateColumn(values)

	//***********************
	// 不要采用下面的方式更新数据，当第一条语句有错误时，将更新所有数据（除非判断错误并抛出异常）
	//***********************
	//conn.DB.Where("id > ", 3).First(&users)
	//conn.DB.Model(&users).Select("null_age", "address").UpdateColumn(values)
}

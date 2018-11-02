package dmloperator

import (
	"database/sql"
	"fmt"
	"gorm_demo/conn"
	"gorm_demo/model"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
)

func InsertOne(idx int) {
	goid := strconv.Itoa(idx)
	fmt.Println(goid)

	var user model.User
	user.Address = "aaa"
	user.CreatedAt = time.Now()
	user.DeptId = sql.NullInt64{Int64: 0, Valid: false}
	user.Email = ""
	user.Name = "wins"
	user.NullAge = sql.NullInt64{Int64: 21, Valid: true}
	user.NullString = ""
	user.Birthday = mysql.NullTime{Time: time.Now(), Valid: false}
	user.GoId = goid
	//user.UpdatedAt ==>会自动更新为time.Now

	user.ID = 0
	db := conn.DB.Create(&user)
	db.Close()
	if db.Error != nil {
		fmt.Println("--------error----")

	}

}

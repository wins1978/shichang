package dml

import (
	"github.com/go-sql-driver/mysql"
	"database/sql"
    "time"
    "strconv"
    "fmt"
    "g.demo/model"
    "g.demo/db"
)

func InsertUser(idx int, deptId int64) {
	var user model.User
	user.Address = "事务测试"
	user.CreatedAt = time.Now()
	user.DeptId = sql.NullInt64{Int64: deptId, Valid: true}
	user.Email = "11777@qq.com"
	user.Name = "wins"
	user.NullAge = sql.NullInt64{Int64: 21, Valid: true}
	user.NullString = ""
	user.Birthday = mysql.NullTime{Time: time.Now(), Valid: false}
    user.GoId=strconv.Itoa(idx)


    fmt.Println(idx)
    user.ID = 0
    
    //time.Sleep(time.Duration(5) * time.Second)
    da :=db.InitDB()

	q := da.Create(&user)
	if q.Error != nil {
        fmt.Println(q.Error)
    }
    defer q.Close()
}

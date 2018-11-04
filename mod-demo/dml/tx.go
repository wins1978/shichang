package dml

import (
	"database/sql"
    "time"
    "fmt"
    "github.com/jinzhu/gorm"
	"g.demo/model"
)

//事务在不同的goroutine中一致性

func TxInsertUser(tx *gorm.DB) {
	var user model.User
	user.Address = "事务测试"
	user.CreatedAt = time.Now()
	//user.DeptId = sql.NullInt64{Int64: 1, Valid: true}
	user.Email = "11777@qq.com"
	user.Name = "wins"
	//user.NullAge = sql.NullInt64{Int64: 21, Valid: true}
	user.NullString = ""
	//user.Birthday = mysql.NullTime{Time: time.Now(), Valid: false}
	//user.GoId = strconv.Itoa(1)
	user.ID = 0

    q :=tx.Create(&user)
    if(q.Error != nil){
        fmt.Println("-------------------TxInsertUser:",q.Error)
        //tx.Rollback()
    }
}

func TxInsertDept(tx *gorm.DB) {
    var flo float64 = 20.23
	var dept model.Department
	dept.Address = ""
	dept.CreatedAt = time.Now()
	dept.DeptName = "企业IT"
	dept.NullFloat = sql.NullFloat64{Float64: flo, Valid: true}
	dept.Tel = "18011112222"
	//dept.UpdatedAt

	q := tx.Create(&dept)
	if q.Error != nil {
        fmt.Println("-------------------TxInsertDept:",q.Error)
        tx.Rollback()
	}
}

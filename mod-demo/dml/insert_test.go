package dml_test

import (
	"database/sql"
	"g.demo/db"
	"g.demo/model"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestMain(m *testing.M){
    m.Run()
}

func TestInsertOne(t *testing.T) {
	var user model.User
	user.Address = "深圳南山"
	user.CreatedAt = time.Now()
	user.DeptId = sql.NullInt64{Int64: 0, Valid: false}
	user.Email = "11777@qq.com"
	user.Name = "wins"
	user.NullAge = sql.NullInt64{Int64: 21, Valid: true}
	user.NullString = ""
	user.Birthday = mysql.NullTime{Time: time.Now(), Valid: false}
	//user.UpdatedAt ==>会自动更新为time.Now

    user.ID = 0
    
    da :=db.InitDB()

	q := da.Create(&user)
	if q.Error != nil {
		t.Error(q.Error)
	}

	var flo float64 = 20.23
	var dept model.Department
	dept.Address = ""
	dept.CreatedAt = time.Now()
	dept.DeptName = "企业IT"
	dept.NullFloat = sql.NullFloat64{Float64: flo, Valid: true}
	dept.Tel = "18011112222"
	//dept.UpdatedAt

	q2 := da.Create(&dept)
	if q2.Error != nil {
		t.Error(q2.Error)
	}
}

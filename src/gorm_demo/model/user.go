package model

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

// User 用户表
type User struct {
	ID         uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string
	NullAge    sql.NullInt64
	Email      string `gorm:"type:varchar(100);not null"`
	NullString string `gorm:"size:255"`   // 设置字段大小为255
	Address    string `gorm:"index:addr"` // 给address字段创建名为addr的索引
	//IgnoreMe     int     `gorm:"-"`               // 忽略本字段
	Birthday mysql.NullTime
	DeptId   sql.NullInt64
}

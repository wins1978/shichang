package temp

import (
"go_web/common"
 "fmt"
//"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/mysql"
"log"
"os"
)

//https://blog.csdn.net/feiwutudou/article/details/82782470
type Test_tb struct {
	ID   int `gorm:"primary_key"` //指定主键并自增
	Address      string
	Phone       string
}

func (Test_tb) TableName() string {
    return "Test_tb"
}

func GetFirstRow() {
	common.InitDB()
	//ttb := Test_tb{}
	//ttb.ID=1

	//data := Test_tb3{}
	
	//SELECT * FROM Test_tb WHERE ID = 1 LIMIT 1;
	//common.Db.First(&ttb,1)
	
	//common.Db.AutoMigrate(&Test_tb{})

	//p1 := Test_tb{ID : 0, Address : "new addr", Phone : "100220033"}
	//common.Db.Create(&p1)

	var tb Test_tb
	common.Db.LogMode(true)
	//common.Db.SetLogger(log.)
	common.Db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	qhuser:=common.Db.First(&tb)
	if err:=qhuser.Error; err!=nil{
        fmt.Print(err)
	}
	
	//common.Db.Debug().Model(&tb).Update("ADDRESS", "NNNN")
	defer common.Db.Close()

	fmt.Print(tb.Phone)
	
}
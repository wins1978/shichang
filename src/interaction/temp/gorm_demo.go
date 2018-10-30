package temp

import (
"go_web/common"
 "fmt"
"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/mysql"
"log"
"os"
)

//http://gorm.io/docs/query.html
//https://blog.csdn.net/feiwutudou/article/details/82782470
type Test_tbStruct struct {
	gorm.ModelStruct
	id   int `gorm:"primary_key;Column:ID"` //指定主键并自增
	ADDRESS      string `gorm:"Column:ADDRESS"`
	phone       string `gorm:"Column:PHONE"`
}

func (Test_tbStruct) TableName() string {
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

	var tb Test_tbStruct
	common.Db.LogMode(true)
	//common.Db.SetLogger(log.)
	common.Db.SetLogger(log.New(os.Stdout, "\r\n", 0))



	q:=common.Db.First(&tb)
	if err:=q.Error; err!=nil{
        fmt.Print(err)
	}

	//var addr string
	rows,_ := common.Db.Model(&Test_tbStruct{}).Where("ID = ?", 1).Rows() // (*sql.Row)
	//rr := common.Db..Where("ID = ?", 1).Select("ADDRESS").Row() // (*sql.Row)
	defer rows.Close()
	
	fmt.Println("==============")
	for rows.Next() {
		//var address string 
		//var phone string
		var rr Test_tbStruct
		//common.Db.ScanRows(rows,user)
		//rows.Scan(&address,&phone)
		common.Db.ScanRows(rows, &rr)
		fmt.Println(rr.phone)
		//fmt.Println(address, " ", phone)
	}

	//convert.go

	fmt.Println("==============")
	
	
	//common.Db.Debug().Model(&tb).Update("ADDRESS", "NNNN")
	defer common.Db.Close()

	
	
	
}
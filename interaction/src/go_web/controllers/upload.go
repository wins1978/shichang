package controllers
/*
1) Upload Exc file to server path
2) Read from Exc plugin
3) Map with Fix DB column
4) Validate column -- not null
5) Show UI with warning

6) Save sourcedata to DB -- EXCEL_IMPORT_DAILY
7) Create Daily Report
8) Create Daily WX Report

EXCEL======
1) Read from Exc --Ignore DATA, params--column mapping
2) Write to Exc -- Ignore DATA
*/

import (
	"github.com/astaxie/beego"
	"fmt"
	"go_web/config"
	"go_web/business/excel"
)

type UploadController struct {
	beego.Controller
}
//https://www.cnblogs.com/goloving/p/8967865.html 上传文件示例
//http://localhost:8080/Upload/UpFile
func (this *UploadController) UpFile() {
	fmt.Println("calling upload")
	f,h,_ := this.GetFile("file")
	fmt.Println(h.Filename)
	fmt.Println(f)

	defer f.Close()
	path := config.UPLOAD_FOLDER +h.Filename
	this.SaveToFile("file",path)

	var colNames = []string{"str1", "str2", "str3", "str4"}

	excReader := new(excel.ExcelReader)
	excReader.ReadFormFile(path,colNames)
	
	this.Ctx.WriteString( "上传成功" )
}
package controllers

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

	excReader := new(excel.ExcelReader)
	excReader.ReadFormFile(path)

	this.Ctx.WriteString( "上传成功" )
}
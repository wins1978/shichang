package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type UploadController struct {
	beego.Controller
}
//https://www.cnblogs.com/goloving/p/8967865.html 上传文件示例
//http://localhost:8080/Upload
func (this *UploadController) UpFile() {
	fmt.Println("calling upload")
	f,h,_ := this.GetFile("file")
	fmt.Println(h.Filename)
	fmt.Println(f)
	this.Ctx.WriteString( "上传成功~！！！！！！！" )
}
package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UploadController struct {
	beego.Controller
}
//https://www.cnblogs.com/goloving/p/8967865.html 上传文件示例
//http://localhost:8080/Upload
func (this *UploadController) UpFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("calling upload")
	s, _ := ioutil.ReadAll(r.Body) //把  body 内容读入字符串 s
    fmt.Fprintf(w, "%s", s)        //在返回页面中显示内容。

	this.Ctx.WriteString( "上传成功~！！！！！！！" )
}
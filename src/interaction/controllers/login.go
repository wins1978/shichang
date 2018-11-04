package controllers

import (
	"fmt"
	"interaction/demo"
	"net/http"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

//http://localhost:8080/Login/GetUserInfo
func (c *LoginController) GetUserInfo() {
	fmt.Println("calling login")
	var da demo.DAA
	da.CreateUser("this is url")

	http.HandleFunc("/hello", hello(da))

	c.Data["Account"] = "wins2"
	c.Ctx.Output.JSON(c.Data["Account"], false, false)
}

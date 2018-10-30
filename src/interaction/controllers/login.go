package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type LoginController struct {
	beego.Controller
}
//http://localhost:8080/Login/GetUserInfo
func (c *LoginController) GetUserInfo() {
	fmt.Println("calling login")
	c.Data["Account"] = "wins"
	c.Ctx.Output.JSON(c.Data["Account"],false,false)
}
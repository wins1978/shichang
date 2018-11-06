package controllers

import (
	"fmt"
	"interaction/demo"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

//http://localhost:8080/Login/GetUserInfo
func (c *LoginController) GetUserInfo() {
	fmt.Println("calling login")

	demo.CreateUser()

	c.Data["Account"] = "wins2"
	c.Ctx.Output.JSON(c.Data["Account"], false, false)
}

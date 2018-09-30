package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) GetUserInfo() {
	c.Data["Account"] = "wins"
	
	c.Ctx.Output.JSON(c.Data["Account"],false,false)
}
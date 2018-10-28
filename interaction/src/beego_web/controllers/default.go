package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego23.me"
	//c.Data["Email"] = "astaxie1@gmail.com"
	c.TplName = "index.tpl"
}


func (c *MainController) GetAll() {
	//c.Data["Website"] = "beego12.me"
	//c.Data["Email"] = "astaxie1@gmail.com"
	c.TplName = "index.tpl"
}

package controllers

import (
	"github.com/astaxie/beego"
	"go_web/temp"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego23.me"
	//c.Data["Email"] = "astaxie1@gmail.com"
	//c.TplName = "index.tpl"
	temp.GetFirstRow()

	c.Ctx.Output.JSON("",false,false)
}


func (c *MainController) GetAll() {
	//c.Data["Website"] = "beego12.me"
	//c.Data["Email"] = "astaxie1@gmail.com"
	c.TplName = "index.tpl"
}

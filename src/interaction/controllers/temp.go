package controllers

import (
	"interaction/temp"

	"github.com/astaxie/beego"
)

type TempController struct {
	beego.Controller
}

func (this *TempController) Get() {
	temp.ReflectDemo()

	this.Ctx.Output.JSON("temp...", false, false)
}

package controllers
import (
	"github.com/astaxie/beego"
	"go_web/temp"
)
type TempController struct {
	beego.Controller
}

func (this *TempController)Get() {
	temp.ReflectDemo()

	this.Ctx.Output.JSON("temp...", false, false)
}


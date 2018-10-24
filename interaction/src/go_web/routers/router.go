package routers

import (
	"github.com/astaxie/beego/plugins/cors"
	"go_web/controllers"
	"github.com/astaxie/beego"
)


func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.LoginController{})
}


package main

import (
	"interaction/common"
	_ "interaction/routers"

	"github.com/astaxie/beego"
	//"go_web/common"
)

func init() {
	// init db env
	common.InitDB()
}

func main() {
	beego.Run()
}

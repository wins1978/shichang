package main

import (
	"github.com/astaxie/beego"
	_ "go_web/routers"
	//"go_web/common"
)




func init() {
	// init db env
	//common.InitDB()
}

func main() {
	beego.Run()
}

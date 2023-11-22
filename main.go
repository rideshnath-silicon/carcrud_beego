package main

import (
	_ "CarCrudDemo/routers"

	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
)

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

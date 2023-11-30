package main

import (
	_ "CarCrudDemo/routers"

	"github.com/astaxie/beego"
	"github.com/beego/beego/orm"
	_ "github.com/lib/pq"
)

func main() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("sqlconn"))

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

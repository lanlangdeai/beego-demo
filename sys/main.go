package main

import (
	_ "sys/routers"
	_"sys/init"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}


package main

import (
	"blog/controllers"
	_ "blog/routers"
	"github.com/astaxie/beego"
)

func main() {

	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

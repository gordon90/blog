package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

func (self *ErrorController) Error404() {
	self.TplName = "error/error404.tpl"
}

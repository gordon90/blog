package controllers

import "github.com/astaxie/beego"

type AboutController struct {
	beego.Controller
}

func (self *AboutController) Get() {
	self.TplName = "home/about.html"
}

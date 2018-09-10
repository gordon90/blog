package controllers

import "github.com/astaxie/beego"

type PhotosController struct {
	beego.Controller
}

func (self *PhotosController) Get() {
	self.TplName = "home/share.html"
}

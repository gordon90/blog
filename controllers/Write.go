package controllers

import (
	"github.com/astaxie/beego"
)

type WriteController struct {
	beego.Controller
}

func (self *WriteController) Get() {

	self.TplName = "write/index.html"
}

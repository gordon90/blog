package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CateController struct {
	beego.Controller
}
type Cate struct {
	ID   int
	Name string
	Nums int
}

func init() {
	orm.RegisterModel(new(Cate))
}

func (self *CateController) Add() {
	self.TplName = "write/cate_add.html"
}

func (self *CateController) AddPost() {
	o := orm.NewOrm()
	sql := new(Cate)

	sql.Name = self.GetString("cate")
	_, err := o.Insert(sql)
	if err == nil {
		self.Redirect("/cate_list", 301)
	} else {
		self.Ctx.WriteString("err")
	}
}

func (self *CateController) CateList() {

	o := orm.NewOrm()

	_, err := o.Raw("select * from cate").Values(&maps)
	if err != nil {
		self.Ctx.WriteString("err")
		panic(self)
	}

	self.Data["cate"] = maps
	self.TplName = "write/cate_list.html"
}

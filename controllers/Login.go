package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// post parameter
type LoginController struct {
	beego.Controller
}

type User struct {
	Id       int
	Username string
	Password string
}

func init() {
	// set default database
	orm.RegisterDataBase("user", "mysql", "root:root@tcp(localhost:3306)/blog?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(User))
}

func (self *LoginController) Get() {
	self.TplName = "write/login.tpl"
}

func (self *LoginController) Post() {

	o := orm.NewOrm()

	user := User{}
	//
	paramUsername := self.GetString("Username")
	paramPassword := self.GetString("Password")

	// 加密
	h := sha1.New()
	h.Write([]byte(paramPassword))
	hash_password := fmt.Sprintf("%x", h.Sum(nil)) // a5e05e6a02500dae95519eea8d0d1b74e0730cf1
	err := o.Raw("select id from user where username=? AND password=?", paramUsername, hash_password).QueryRow(&user)

	if err == nil {
		self.Redirect("/write", 301)
		//self.Header().Set("Location", "url") w.WriteHeader(301)
	} else {
		self.Redirect("/login", 301)
	}
}

func (self *LoginController) Veri() {
	self.Ctx.WriteString("aaa")
}

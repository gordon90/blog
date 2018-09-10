package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Contact struct {
	ID          int
	Name        string
	Email       string
	Content     string
	Create_time int64
	//StringTime  string
}

type GbookController struct {
	beego.Controller
}

func init() {
	orm.RegisterModel(new(Contact))
}

func (self *GbookController) Get() {

	pageParam := self.GetString("page")
	o := orm.NewOrm()
	var page int
	var pagePointer *int = &page
	if pageParam == "" {
		*pagePointer = 1
	} else {
		*pagePointer, _ = strconv.Atoi(pageParam)
	}

	self.Data["current_page"] = page
	o.Raw("select * from contact order by create_time limit ?, 6", (page-1)*6).Values(&maps)
	self.Data["contacts"] = maps

	var pages []orm.Params
	nums, _ := o.Raw("select * from contact").Values(&pages)
	self.Data["total_page"] = (nums + 6 - 1) / 6

	for i, _ := range maps {
		maps[i]["create_time"] = time.Unix(TimeFormat(maps[i]["create_time"]), 0).Format("2006-01-02")
	}

	self.TplName = "home/gbook.html"
}

func TimeFormat(time interface{}) int64 {
	timeString, _ := time.(string)
	timeInt64, _ := strconv.ParseInt(timeString, 10, 64)
	return timeInt64
}

func (self *GbookController) Post() {
	o := orm.NewOrm()
	name := self.GetString("name")
	email := self.GetString("email")
	content := self.GetString("content")

	contact := Contact{}
	contact.Name = name
	contact.Email = email
	contact.Content = content
	contact.Create_time = time.Now().Unix()

	_, err := o.Insert(&contact)
	fmt.Println(err)
	if err == nil {
		//self.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": imgpath[6]}
		self.Data["json"] = map[string]interface{}{
			"code": 200,
		}
	} else {
		self.Data["json"] = map[string]interface{}{
			"code": 201,
		}
	}
	self.ServeJSON()

}

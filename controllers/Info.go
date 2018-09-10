package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type InfoController struct {
	beego.Controller
}

type Read_record struct {
	Id  int
	Aid int
	Ip  string
}

func init() {
	orm.RegisterModel(new(Read_record))
}
func (self *InfoController) Get() {
	id := self.GetString("id")
	o := orm.NewOrm()

	var article Article

	if id == "" {
		o.Raw("select id,title,content,read_nums,create_time from article order by create_time DESC limit 1").QueryRow(&article)
		readCount(self, article.Id)
	} else {
		o.Raw("select id,title,content,read_nums,create_time from article where id=?", id).QueryRow(&article)
		readCount(self, article.Id)
	}
	self.Data["time"] = time.Unix(article.Create_time, 0).Format("2006-01-02")
	self.Data["title"] = article.Title
	self.Data["content"] = article.Content
	self.Data["read_nums"] = article.Read_nums

	//  分类
	var cate []orm.Params
	o.Raw("select id,name,nums from cate").Values(&cate)
	self.Data["cate"] = cate

	// 上一篇 下一篇
	var next Article
	var previous Article
	o.Raw("select id,title from article where id>? limit 1", id).QueryRow(&next)
	o.Raw("select id,title from article where id<? limit 1", id).QueryRow(&previous)

	if next.Id != 0 {
		self.Data["next_true"] = true
	}
	self.Data["next_title"] = next.Title
	self.Data["next_id"] = next.Id

	if previous.Id != 0 {
		self.Data["previous_true"] = true
	}

	self.Data["previous_title"] = previous.Title
	self.Data["previous_id"] = previous.Id
	self.TplName = "home/info.html"
}

func (self *InfoController) Digg() {
	id := self.GetString("id")
	self.Ctx.WriteString(id)
}

func (self *InfoController) Vague() {
	input := self.GetString("keyword")
	var articles []orm.Params
	o := orm.NewOrm()
	input = "%" + input + "%"
	nums, _ := o.Raw("select id, title from article where title like ?", input).Values(&articles)
	fmt.Println(articles)
	if nums > 0 {
		self.Data["json"] = map[string]interface{}{
			"code": 200,
			"data": articles,
		}
	} else {
		self.Data["json"] = map[string]interface{}{
			"code": 202,
		}
	}
	self.ServeJSON()
}

// 阅读量
func readCount(self *InfoController, id int) {

	ip := self.Ctx.Input.IP()

	o := orm.NewOrm()
	nums, _ := o.Raw("select id from read_record where aid=? AND ip=?", id, ip).Values(&maps)

	if nums == 0 {
		readRecod := new(Read_record)
		readRecod.Ip = ip
		readRecod.Aid = id
		_, err := o.Insert(readRecod)
		if err != nil {
			beego.Error(err)
		}

		var article Article
		o.Raw("select read_nums from article where id=?", id).QueryRow(&article)
		article.Read_nums += 1
		res, err := o.Raw("update article set read_nums=? where id=?", article.Read_nums, id).Exec()
		if err == nil {
			res.RowsAffected()
		}
	}
	return
}

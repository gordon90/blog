package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type value struct {
	Cate_id     int
	Title       string
	Keyword     string
	Content     string
	Author      string
	create_time int64
}

type ArticleController struct {
	beego.Controller
}

func (self *ArticleController) Add() {

	o := orm.NewOrm()
	_, err := o.Raw("select id, name from cate").Values(&maps)
	if err != nil {
		self.Ctx.WriteString("err")
	}
	self.Data["cate"] = maps
	fmt.Println(maps)
	self.TplName = "write/article_add.html"
}

func (self *ArticleController) AddPost() {

	o := orm.NewOrm()

	param := value{}
	article := new(Article)

	// Must
	if err := self.ParseForm(&param); err != nil {
		self.Ctx.WriteString("err")
	}

	article.Cate_id = param.Cate_id
	article.Title = param.Title
	article.Content = param.Content
	article.Author = param.Author
	article.Keyword = param.Keyword
	article.Create_time = time.Now().Unix()

	_, err := o.Insert(article)
	if err == nil {
		var cate Cate
		err := o.Raw("select nums from cate where id=?", param.Cate_id).QueryRow(&cate)
		if err != nil {
			self.Ctx.WriteString("err")
		}

		cate.Nums = cate.Nums + 1
		// 类别文章数量
		res, err := o.Raw("update cate set nums=? where id=?", cate.Nums, param.Cate_id).Exec()
		if err != nil {
			self.Ctx.WriteString("err")
		}
		res.RowsAffected()

		self.Data["json"] = map[string]interface{}{
			"code": 200,
			"msg":  "success",
		}
		self.ServeJSON()
	} else {
		self.Data["json"] = map[string]interface{}{
			"code": 201,
			"msg":  "falure",
		}
		self.ServeJSON()
	}
}

func (self *ArticleController) List() {

	o := orm.NewOrm()
	_, err := o.Raw("select id,title,content,create_time from article order by create_time limit 10").Values(&maps)

	if err == nil {
		for i, _ := range maps {
			maps[i]["content"] = TrimHtml(fmt.Sprint(maps[i]["content"]))
		}
	}
	self.Data["data"] = maps
	var cates []orm.Params
	o.Raw("Selcet id, name from cate").Values(&cates)
	self.Data["cate"] = cates

	self.TplName = "write/article_list.html"
}

func (self *ArticleController) ArticleDel() {
	param_id := self.Ctx.Input.Param(":id")

	id, _ := strconv.Atoi(param_id) // 类型转换
	o := orm.NewOrm()

	article := new(Article)
	article.Id = id
	o.Raw("select cate_id from article where id=?", param_id).QueryRow(&article)
	_, err := o.Delete(article)

	if err == nil {
		var cate Cate
		o.Raw("select nums from cate where id=?", article.Cate_id).QueryRow(&cate)
		cate.Nums = cate.Nums - 1
		res, _ := o.Raw("update cate set nums=? where id=?", cate.Nums, article.Cate_id).Exec()
		res.RowsAffected()
		fmt.Println(cate.Nums)
		self.Redirect("/article_list", 301)
	} else {
		self.Ctx.WriteString("err")
	}
}

func (c *ArticleController) UploadImg() {
	// 获取文件
	_, h, err := c.GetFile("file")

	if err != nil {
		beego.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix

	path := "static/img/"

	year, month, _ := time.Now().Date()
	subdirectory := strconv.Itoa(year) + month.String() + "/"
	os.Mkdir(path+subdirectory, 0777)
	fileUploadPath := path + subdirectory

	err = c.SaveToFile("file", fileUploadPath+newname)

	if err != nil {
		beego.Error(err)
	}

	url := "http://localhost:5050"
	c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": url + "/" + fileUploadPath + newname}
	c.ServeJSON()
}

func (self *ArticleController) Del() {
	imgUrl := self.GetString("imgUrl")

	imgpath := strings.Split(imgUrl, "/")
	err := os.Remove("static/img/" + imgpath[5] + "/" + imgpath[6])

	if err == nil {
		self.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": imgpath[6]}
	} else {
		self.Data["json"] = map[string]interface{}{"state": "ERROR", "link": ""}
	}

	self.ServeJSON()
}

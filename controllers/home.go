package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"regexp"
	"strconv"
	"strings"
	"time"
)

type HomeController struct {
	beego.Controller
}

type Article struct {
	Id          int
	Cate_id     int
	Title       string
	Keyword     string
	Content     string
	Author      string
	Read_nums   int
	Create_time int64
}

var maps []orm.Params

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/blog?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(Article))
}

func (self *HomeController) Get() {

	var o = orm.NewOrm()
	cateParam := self.GetString("cate_id")
	pageParam := self.GetString("page")

	keyword := self.GetString("keyword")

	// 分页
	var page int
	var pagePointer *int = &page
	if pageParam == "" {
		*pagePointer = 1
	} else {
		*pagePointer, _ = strconv.Atoi(pageParam)
	}
	self.Data["current_page"] = page
	cate_id, _ := strconv.Atoi(cateParam)
	keyword = "%" + keyword + "%"
	if cateParam == "" {
		if keyword == "" {
			o.Raw("select id,title,create_time,content from article order by create_time DESC limit ?,10", (page-1)*10).Values(&maps)
		} else {
			o.Raw("select id,title,create_time,content from article where title like ? order by create_time DESC limit ?,10", keyword, (page-1)*10).Values(&maps)
		}
		var pages []orm.Params
		num, _ := o.Raw("select id from article").Values(&pages)
		self.Data["total_page"] = (num + 10 - 1) / 10
	} else {
		o.Raw("select id,title,create_time,content from article where cate_id=? order by create_time DESC limit ?,10", cate_id, (page-1)*10).Values(&maps)
		var pages []orm.Params
		num, _ := o.Raw("select * from article where cate_id=?", cate_id).Values(&pages)
		self.Data["total_page"] = (num + 10 - 1) / 10
	}

	for key, _ := range maps {
		maps[key]["content"] = TrimHtml(fmt.Sprintf("%s", maps[key]["content"]))
		maps[key]["create_time"] = time.Unix(TimeFormat(maps[key]["create_time"]), 0).Format("2006-01-02")

	}
	self.Data["content"] = maps

	// 分类
	var cates []orm.Params
	o.Raw("select id,name,nums from cate").Values(&cates)
	self.Data["cate"] = cates

	self.TplName = "home/index.html"

}

func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	return strings.TrimSpace(src)
}

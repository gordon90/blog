package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	/* Home Router*/
	beego.Router("/", &controllers.HomeController{})
	// information
	beego.Router("/info", &controllers.InfoController{})
	// contact
	beego.Router("/gbook", &controllers.GbookController{})

	/* Admin Router*/
	beego.Router("/write", &controllers.WriteController{})
	// article
	beego.Router("/article_add", &controllers.ArticleController{}, "Get:Add")
	beego.Router("/article_post", &controllers.ArticleController{}, "Post:AddPost")
	beego.Router("/article_list", &controllers.ArticleController{}, "Get:List")
	beego.Router("/article_delete/:id:int", &controllers.ArticleController{}, "Get:ArticleDel")
	// cate
	beego.Router("/cate_add", &controllers.CateController{}, "Get:Add")
	beego.Router("/cate_post", &controllers.CateController{}, "Post:AddPost")
	beego.Router("/cate_list", &controllers.CateController{}, "Get:CateList")

	// image upload
	beego.Router("/uploadimg", &controllers.ArticleController{}, "Post:UploadImg")
	beego.Router("/delete", &controllers.ArticleController{}, "Get:Del")

	// 未写
	beego.Router("/vauge", &controllers.InfoController{}, "Post:Vague")
	beego.Router("/digg", &controllers.InfoController{}, "Get:Digg")
	beego.Router("/about", &controllers.AboutController{})
	beego.Router("/share", &controllers.PhotosController{})

	// login Router
	beego.Router("/login", &controllers.LoginController{})
}

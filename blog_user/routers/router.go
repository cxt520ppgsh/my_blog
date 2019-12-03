package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"go_vue_blog/blog_user/controllers"
)

func init() {
	// 这段代码放在router.go文件的init()的开头
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"http://39.100.119.126:8000",
			"http://39.100.119.126:4000",
			"http://" + beego.AppConfig.String("front_end_domain") + ":" + beego.AppConfig.String("front_end_port")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.Router("/", &controllers.MainController{})
	beego.Router("/articleDetail", &controllers.ArticleDetailController{})
	beego.Router("/articleList", &controllers.ArticleListController{})

}

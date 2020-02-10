package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"go_vue_blog/blog_user/controllers"
	"go_vue_blog/blog_user/controllers/utils"
	"os"
)

func init() {
	// 这段代码放在router.go文件的init()的开头
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{
			"http://39.100.119.126:8000",
			"http://39.100.119.126:4000",
			"http://127.0.0.1:4000",
			"http://127.0.0.1:8000",
			"http://0.0.0.0:8000",
			"http://0.0.0.0:4000",
			"http://" + beego.AppConfig.String("front_end_domain") + ":" + beego.AppConfig.String("front_end_port")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.Router("/", &controllers.MainController{})
	beego.Router("/articleDetail", &controllers.ArticleDetailController{})
	beego.Router("/articleList", &controllers.ArticleListController{})

	test()
}

//const SSH_USERNAME = "chenxutang"
//const SSH_PASSWORD = "fly098321"
//const SSH_PROJECT_PATH = "project/010"
//const SSH_IP = "172.168.1.187"
//const SSH_PORT = 22

const SSH_USERNAME = "chenxutang"
const SSH_PASSWORD = "T373683458"
const SSH_PROJECT_PATH = "project"
const SSH_IP = "39.100.119.126"
const SSH_PORT = 22

func test()  {
	cli := utils.New(SSH_IP, SSH_USERNAME, SSH_PASSWORD, SSH_PORT)
	output, err := cli.Run("cd " + SSH_PROJECT_PATH+";pwd")
	fmt.Printf("%v\n%v", output, err)
	cli.RunTerminal("top", os.Stdout, os.Stdin)
}

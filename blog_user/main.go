package main

import (
	"go_vue_blog/blog_user/controllers/utils"
	_ "go_vue_blog/blog_user/routers"
	"github.com/astaxie/beego"
)

func main() {
	go utils.StartSocket()
	beego.Run()
}


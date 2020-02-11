package controllers

import (
	_ "bytes"
	"github.com/astaxie/beego"
	"go_vue_blog/blog_user/controllers/utils"
	_ "log"
	_ "os/exec"
	_ "strconv"
	_ "strings"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"

	utils.TestBuild()
}

package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"go_vue_blog/blog_user/controllers/utils"
	_"go_vue_blog/blog_user/models"
)

type SSHBuildController struct {
	beego.Controller
}


func (this *SSHBuildController) Post() {

	result := make(map[string]interface{})

	result["code"] = "000"
	result["msg"] = "success"
	result["data"] = "data"

	bytes, err := json.Marshal(result)
	if err != nil{
		fmt.Println(err)
	}

	this.Data["json"] = string(bytes)
	this.ServeJSON()
	utils.TestBuild()
}

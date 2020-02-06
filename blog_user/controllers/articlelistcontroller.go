package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"go_vue_blog/blog_user/models"
)

type ArticleListController struct { 
	beego.Controller
}


func (this *ArticleListController) Post() {
	models.QueryMDArticleList()

	result := make(map[string]interface{})
	
	result["code"] = "000"
	result["msg"] = "success"
	//var articleList []models.ArticleResult = models.QueryArticleList() 读数据库
	var articleList []models.MDArticleResult = models.QueryMDArticleList() //读static文件夹的MD
	result["data"] = articleList

	bytes, err := json.Marshal(result)
	if err != nil{
		fmt.Println(err)
	}

	this.Data["json"] = string(bytes)
	this.ServeJSON()
}

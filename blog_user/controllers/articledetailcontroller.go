package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_"go_vue_blog/blog_user/controllers/utils"
	"go_vue_blog/blog_user/models"
)

type ArticleDetailController struct {
	beego.Controller
}

func (this *ArticleDetailController) Post() {
	result := make(map[string]interface{})

	var ob models.ArticleDetail
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)

	//var articleDetail models.ArticleDetailResult = models.QueryArticleDetail(ob.ArticleId) //使用数据库查ID
	var articleDetail models.MDArticleDetailResult = models.QueryMDArticleDetail(ob.ArticleId) //查static目录
	fmt.Println("request", string(this.Ctx.Input.RequestBody))
	result["code"] = "000"
	result["msg"] = "success"
	result["data"] = articleDetail
	bytes, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	this.Data["json"] = string(bytes)
	this.ServeJSON()
}



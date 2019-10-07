package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"go_vue_blog/blog_user/models"
)

type ArticleDetailController struct {
	beego.Controller
}

func (this *ArticleDetailController) Post() {
	result := make(map[string]interface{})

	var ob models.ArticleDetail
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)

	var articleDetail models.ArticleDetailResult = models.QueryArticleDetail(ob.ArticleId)
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

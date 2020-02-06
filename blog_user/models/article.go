package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "io/ioutil"
	_ "os"
	"strings"
)

type Article struct{
	ArticleId string `json:"articleId"` //首字母大写表示public，小写表示private，添加tag，使其输出时变成小写
	Title string `json:"title"`
	Date string `json:"date"`
	Gist string `json:"gist"`
	Labels string `json:"labels"`
}

type ArticleResult struct{
	ArticleId string `json:"articleId"` //首字母大写表示public，小写表示private，添加tag，使其输出时变成小写
	Title string `json:"title"`
	Date string `json:"date"`
	Gist string `json:"gist"`
	Labels []string `json:"labels"`
}


//查询文章列表
func QueryArticleList() []ArticleResult{
	var dbhost string = beego.AppConfig.String("dbhost")
	var dbport string = beego.AppConfig.String("dbport")
	var dbuser string = beego.AppConfig.String("dbuser")
	var dbpassword string = beego.AppConfig.String("dbpassword")
	var dbname string = beego.AppConfig.String("dbname")
	var dbcharset string = beego.AppConfig.String("dbcharset")

	db, err := sql.Open("mysql", dbuser+":"+dbpassword+"@tcp("+dbhost+":"+dbport+")/"+dbname+"?"+dbcharset)
	if(err != nil){
		return nil
	}

	var sql string = "SELECT articleId,title,date,gist,labels FROM article"
	fmt.Println(sql)
	rows, err := db.Query(sql)
	if(err != nil){
		panic(err)
		return nil
	}

	var articleResults []ArticleResult
	for rows.Next() {

		var article Article
		var articleResult ArticleResult
		rows.Columns()
		err = rows.Scan(&article.ArticleId, &article.Title, &article.Date, &article.Gist, &article.Labels)
		if(err != nil){
			return nil
		}
		var pureLabelString string = strings.Replace(article.Labels,"\u0000","",-1)
		var labels[] string = strings.Split(pureLabelString,",")
		articleResult.Labels = labels
		articleResult.ArticleId = article.ArticleId
		articleResult.Date = article.Date
		articleResult.Gist = article.Gist
		articleResult.Title = article.Title

		articleResults = append(articleResults, articleResult)
	}
	rows.Close()
	db.Close()
	return articleResults
}





type ArticleDetail struct{
	ArticleId string `json:"articleId"` //首字母大写表示public，小写表示private，添加tag，使其输出时变成小写
	Title string `json:"title"`
	Date string `json:"date"`
	Content string `json:"content"`
	Gist string `json:"gist"`
	Labels string `json:"labels"`
}

//传到客户端的参数中，labels需要string数组
type ArticleDetailResult struct{
	ArticleId string `json:"articleId"` //首字母大写表示public，小写表示private，添加tag，使其输出时变成小写
	Title string `json:"title"`
	Date string `json:"date"`
	Content string `json:"content"`
	Gist string `json:"gist"`
	Labels []string `json:"labels"`
}


//查询文章详情
func QueryArticleDetail(articleId string) ArticleDetailResult{
	var dbhost string = beego.AppConfig.String("dbhost")
	var dbport string = beego.AppConfig.String("dbport")
	var dbuser string = beego.AppConfig.String("dbuser")
	var dbpassword string = beego.AppConfig.String("dbpassword")
	var dbname string = beego.AppConfig.String("dbname")
	var dbcharset string = beego.AppConfig.String("dbcharset")

	var articleResult ArticleDetailResult
	var article ArticleDetail
	db, err := sql.Open("mysql", dbuser+":"+dbpassword+"@tcp("+dbhost+":"+dbport+")/"+dbname+"?"+dbcharset)
	if(err != nil){
		panic(err)
		return articleResult
	}

	var sql string = "SELECT articleId,title,content,date,gist,labels FROM article where articleId = \"" + articleId + "\""
	fmt.Println(sql)
	rows, err := db.Query(sql)
	if(err != nil){
		panic(err)
		return articleResult
	}

	for rows.Next() {

		rows.Columns()
		err = rows.Scan(&article.ArticleId, &article.Title,&article.Content, &article.Date, &article.Gist, &article.Labels)
		if(err != nil){
			return articleResult
		}

	}
	var content string = article.Content
	content = strings.Replace(content,`\"` , `"`, -1)
	content = strings.Replace(content,`\'` , `'`, -1)
	content = strings.Replace(content,`\\\` , `\`, -1)
	article.Content = content

	//解决乱码
	var pureLabelString string = strings.Replace(article.Labels,"\u0000","",-1)
	var labels[] string = strings.Split(pureLabelString,",")
	articleResult.Labels = labels
	articleResult.ArticleId = article.ArticleId
	articleResult.Content = article.Content
	articleResult.Date = article.Date
	articleResult.Gist = article.Gist
	articleResult.Title = article.Title

	rows.Close()
	db.Close()
	return articleResult
}



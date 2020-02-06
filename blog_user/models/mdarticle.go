package models

import (
	"bufio"
	_"database/sql"
	"fmt"
	_"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"io/ioutil"
	_ "io/ioutil"
	"os"
	_ "os"
	"strings"
)
const mdDir = "./static/mdarticles"

type MDArticle struct {
	ArticleId string `json:"articleId"` //首字母大写表示public，小写表示private，添加tag，使其输出时变成小写
	Title     string `json:"title"`
	Date      string `json:"date"`
	Gist      string `json:"gist"`
	Labels    string `json:"labels"`
}

type MDArticleResult struct {
	ArticleId string   `json:"articleId"` //首字母大写表示public，小写表示private，添加tag，使其输出时变成小写
	Title     string   `json:"title"`
	Date      string   `json:"date"`
	Gist      string   `json:"gist"`
	Labels    []string `json:"labels"`
}

type MDArticleDetail struct {
	ArticleId string `json:"articleId"` //首字母大写表示public，小写表示private，添加tag，使其输出时变成小写
	Title     string `json:"title"`
	Date      string `json:"date"`
	Content   string `json:"content"`
	Gist      string `json:"gist"`
	Labels    string `json:"labels"`
}

//传到客户端的参数中，labels需要string数组
type MDArticleDetailResult struct {
	ArticleId string   `json:"articleId"` //首字母大写表示public，小写表示private，添加tag，使其输出时变成小写
	Title     string   `json:"title"`
	Date      string   `json:"date"`
	Content   string   `json:"content"`
	Gist      string   `json:"gist"`
	Labels    []string `json:"labels"`
}

//查询文章列表
func QueryMDArticleList() []MDArticleResult {
	var articleResults []MDArticleResult
	mdFileList, _ := getAllFile(mdDir)
	for _, filename := range mdFileList {
		var articleResult MDArticleResult
		articleResult.Labels = []string{"无"}
		articleResult.ArticleId = filename
		//articleResult.Content = readFileToContent(mdDir + "/" + filename)
		articleResult.Date = strings.Split(strings.Split(filename, "_")[1],".md")[0]
		articleResult.Gist = readFileToGist(mdDir + "/" + filename)
		articleResult.Title = strings.Split(filename, "_")[0]
		articleResults = append(articleResults, articleResult)
	}
	return articleResults
}

//查询文章详情
func QueryMDArticleDetail(articlePath string) MDArticleDetailResult {
	var articleResult MDArticleDetailResult
	filename := strings.Split(articlePath,mdDir+"/")[1]
	articleResult.Labels = []string{"无"}
	articleResult.Content = readFileToContent(articlePath)
	articleResult.Date = strings.Split(strings.Split(filename, "_")[1],".md")[0]
	articleResult.Gist = readFileToGist(mdDir + "/" + filename)
	articleResult.Title = strings.Split(filename, "_")[0]
	return articleResult
}

func getAllFile(pathname string) ([]string, error) {
	var fileList []string
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			getAllFile(pathname + fi.Name() + "\\")
		} else {
			//fmt.Println(fi.Name())
			fileList = append(fileList, fi.Name())
		}
	}
	return fileList, err
}

func readFileToContent(path string) string {
	var content string
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return ""
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		content += line
		if err != nil {
			if err == io.EOF {
				//fmt.Println("File read ok!")
				break
			} else {
				//fmt.Println("Read file error!", err)
				return ""
			}
		}
	}
	return content
}

func readFileToGist(path string) string {
	var content string
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return ""
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	lineCount := 0;
	for {
		lineCount ++
		if lineCount == 3 {
			break
		}
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		content += line
		if err != nil {
			if err == io.EOF {
				//fmt.Println("File read ok!")
				break
			} else {
				//fmt.Println("Read file error!", err)
				return ""
			}
		}
	}
	return content
}

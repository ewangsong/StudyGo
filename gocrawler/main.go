package main

import (
	"github.com/PuerkitoBio/goquery"
	"gocrawler/model"
	"gocrawler/parse"
	"log"
	"net/http"
)

// 分析页面，获取分页
// 分析页面，循环所有页面的电影信息
// 爬去的电影信息入库

//构建请求客户端
var client = &http.Client{}

//构建访问客户端
func httpclient(url string) (doc *goquery.Document) {
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
		return
	}
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	//	request.Header.Set("Accept-Encoding","gzip, deflate, br")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	request.Header.Set("Cache-Control", "max-age=0")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4251.0 Safari/537.36 Edg/87.0.630.0")

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return doc

}
func main() {
	//测试连接数据库
	err := model.InitDB()
	if err != nil {
		log.Fatal("连接数据库失败：err", err)
	}

	url := "https://movie.douban.com/top250"
	doc := httpclient(url)
	pages := parse.ParsePages(doc)

	for _, urls := range pages {
		newurl := url + urls.Url
		doc := httpclient(newurl)
		movies := parse.ParseMovies(doc)
		model.SaveToSql(movies)
	}

}

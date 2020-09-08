package parse

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

// 页面信息
type Page struct {
	Page int
	Url  string
}

// 电影信息
type Movie struct {
	Name  string //名字
	Other string //其他信息
	Year  string //年份
	Area  string //地区
	Tag   string //类型
	Star  string // 评分
	Quote string //评价

}

// 获取所有分页
func ParsePages(doc *goquery.Document) (pages []Page) {
	pages = append(pages, Page{Page: 1, Url: ""})
	doc.Find("#content > div > div.article > div.paginator > a").Each(func(i int, s *goquery.Selection) {
		page, _ := strconv.Atoi(s.Text())
		url, _ := s.Attr("href")
		pages = append(pages, Page{
			Page: page,
			Url:  url,
		})
	})

	return pages
}

// 分析电影信息
func ParseMovies(doc *goquery.Document) (movies []Movie) {
	doc.Find("#content > div > div.article > ol > li").Each(func(i int, s *goquery.Selection) {
		name := s.Find(".title").Eq(0).Text() //名字
		movieinfo := s.Find(".bd p").Eq(0).Text()
		movieinfo = strings.TrimSpace(movieinfo)
		tmps := strings.Split(movieinfo, "\n") //电影信息切片
		other := strings.TrimSpace(tmps[0])

		desc := strings.Split(strings.TrimSpace(string(tmps[1])), "/")
		year := strings.TrimSpace(desc[0])
		area := strings.TrimSpace(desc[1])
		tag := strings.TrimSpace(desc[2])
		star := s.Find(".bd .star .rating_num").Text() //评分
		quote := s.Find(".bd .quote .inq").Text()      //评价

		movies = append(movies, Movie{
			Name:  name,
			Other: other,
			Year:  year,
			Area:  area,
			Tag:   tag,
			Star:  star,
			Quote: quote,
		})
	})

	return movies
}

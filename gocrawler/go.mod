module gocrawler

go 1.14

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/andybalholm/cascadia v1.2.0 // indirect
	gocrawler/model v0.0.0
	gocrawler/parse v0.0.0
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
)

replace (
	gocrawler/model => ./model
	gocrawler/parse => ./parse
)

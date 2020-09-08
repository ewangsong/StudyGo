package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gocrawler/parse"
	"log"
)

var db *sql.DB

func InitDB() (err error) {
	dsn := "root:123456@tcp(192.168.220.138:3306)/test"
	db, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("连接数据库失败err：", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("连接数据库失败err：", err)
	}
	return
}

// SaveToSql 从切片中读取然后存到数据库中
func SaveToSql(m []parse.Movie) {
	sbstr := "insert into sp_douban_movie(name,other,year,area,tag,star,quote) values (?,?,?,?,?,?,?)"

	for _, movice := range m {
		ret, err := db.Exec(sbstr, movice.Name, movice.Other, movice.Year, movice.Area, movice.Tag, movice.Star, movice.Quote)

		if err != nil {
			fmt.Println("插入数据库错误", err, movice)
		}

		id, err := ret.LastInsertId()
		if err != nil {
			fmt.Println(id, err)
		}

	}

}

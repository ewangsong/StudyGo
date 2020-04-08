package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//定义一个全局的db
var db *sql.DB

type user struct {
	Id         int
	TypeId     int
	Name       string
	Password   string
	RealName   string
	CreateTime string
}

//初始化数据库连接函数
func initDB() (err error) {
	dbinfo := "root:test@tcp(172.17.0.40:3306)/lanradius"
	//open 函数只会验证参数格式是否正确不会真正的连接数据库
	db, err = sql.Open("mysql", dbinfo)

	if err != nil {
		return err
	}
	//尝试与数据库进行连接
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("初始化数据库失败", err)
		return
	}
	//查询单行数据
	queryRowFunc()
	//查询多行数据
	queryRowsFunc()
}

//单行查询
func queryRowFunc() {
	sqlStr := "select * from ws_users where id =?"

	var u user
	//
	err := db.QueryRow(sqlStr, 35).Scan(&u.Id, &u.TypeId, &u.Name, &u.Password, &u.RealName, &u.CreateTime)

	if err != nil {
		fmt.Println("查询失败", sqlStr, err)
	}

	fmt.Println(u)

}

//多行查询
func queryRowsFunc() {
	sqlStr := "select * from ws_users where id <?"
	rows, err := db.Query(sqlStr, 20)
	if err != nil {
		fmt.Println("查询失败", sqlStr, err)
		return
	}
	//关闭rows
	defer rows.Close()

	var u user
	var users []user
	for rows.Next() {
		err := rows.Scan(&u.Id, &u.TypeId, &u.Name, &u.Password, &u.RealName, &u.CreateTime)

		if err != nil {
			fmt.Println("scan failed err:", err)
			return
		}
		fmt.Println(u)

		users = append(users, u)

	}
	fmt.Println(users)
}

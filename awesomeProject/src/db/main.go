package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // init(

)

var db *sql.DB

func initDB() (err error) {
	// // 连数据库
	dsn := "root:root@tcp(localhost:3306)/im"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	return
}

type user struct {
	id int
	groupName string
	userId string
}

func query () {

}

func insert() {

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("init db error", err)
	}
	fmt.Println("数据库连接成功")
	var result user
	sqlStr := "select * from group;"
	rowObj := db.QueryRow(sqlStr)
	rowObj.Scan()
	// fmt.Printf("result", result.id, result.groupName, result.userId)
	fmt.Println(result)
}

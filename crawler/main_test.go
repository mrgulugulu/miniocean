package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //注册这个驱动
	"github.com/jinzhu/gorm"
	"testing"
)
type News struct {
	id int `gorm:"AUTO_INCREMENT"`
	news_title string
	news_date string
	news_content string
	news_url string
	news_source string
}
func TestCheckIdentical(t *testing.T) {
	title1 := "萌趣十足！我国首个海龟野化基地迎来龟宝宝孵化季节"
	db, _ := sql.Open("mysql", "root:123456@/hellosea?charset=utf8")
	stmt, err := db.Query(fmt.Sprintf("select id from sea_news_todayhot_v2 where news_title = \"%s\"", title1))
	defer db.Close()
	if err != nil {
		panic(err)
	}
	for stmt.Next(){
		var id int
		err := stmt.Scan(&id)
		if err != nil {
			panic(err)
		}
		fmt.Println(id)
	}
	title2 := "命令是一个按照一定的约定和组织来测试代码的程序"
	stmt, _ = db.Query(fmt.Sprintf("select * from sea_news_todayhot_v2 where news_title = \"%s\"", title2))
	if !stmt.Next() {
		fmt.Println("没有对应数据")
	}
}

func TestExec(t *testing.T) {

	//test := sea_news_todayhot_v2{news_title: "aaaa", news}
	db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&News{})
	test := News{news_title:"aaaa"}
	db.Create(&test)
	var test1 News
	db.Where("news_title=?", "aaaa").First(&test1)
	fmt.Println(test1)
	//
	//var test1 sea_news_todayhot_v2
	//db.Find(&test1, "news_title=?", "aaa")
	//fmt.Println(test1)
	//db.Delete(&test1)
}
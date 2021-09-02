package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strings"
)

var (
	driverName = "mysql"
	user = "root"
	passWord = "123456"
	dataBase = "hellosea"
)

func check(err error) {
	if err != nil{
		panic(err)
	}
}

type Passage struct {
	Title,
	Time,
	Content,
	Url,
	Source string
}

func main() {
	//fetch the structured data from mysql
	db, err := sql.Open(driverName, fmt.Sprintf("%s:%s@/%s?charset=utf8", user, passWord, dataBase))
	check(err)
	rows, err := db.Query("select news_title, news_date, news_content, news_web_url, news_source from sea_news_todayhot_v2 limit 10")
	check(err)
	defer db.Close()

	passages := []Passage{}
	for rows.Next(){
		var(
			news_title,
			news_date,
			news_content,
			news_web_url,
			news_source string
		)
		passage := new(Passage)
		err = rows.Scan(&news_title, &news_date, &news_content, &news_web_url, &news_source)
		check(err)
		passage.Title = news_title
		passage.Time = news_date
		passage.Url = news_web_url
		passage.Content = news_content
		passage.Source = news_source
		passages = append(passages, *passage)
		//fmt.Println(news_title, news_date, news_content, news_web_url, news_source)
	}
	//establish router to show data with json format
	r := gin.Default()
	r.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, passages)
	})
	r.GET("/exact", func(c *gin.Context) {
		title := c.Query("title")
		rows, err := db.Query(fmt.Sprintf("SELECT news_content FROM sea_news_todayhot_v2 WHERE news_title= \"%s\"", title))
		if err != nil {
			for rows.Next() {//一定要next()才能将内容读出来
				var content string
				_ = rows.Scan(&content)
				c.JSON(http.StatusOK, content)
			}
		} else {
			c.String(http.StatusOK, "found nothing")
		}

	})
	r.GET("/dim", func(c *gin.Context) {
		title := c.Query("title")
		rows, err := db.Query("SELECT news_title, news_content FROM sea_news_todayhot_v2")
		if err != nil {
			for rows.Next(){
				var tit, cont string
				_ = rows.Scan(&tit, &cont)
				if strings.Contains(tit, title) {
					c.JSON(http.StatusOK, cont)
				}
			}
		} else {
			c.String(http.StatusOK, "found nothing")
		}
	})
	r.Run()
}

package main

import (
	"crawler/tools"
	_ "github.com/go-sql-driver/mysql" //注册这个驱动
)
//文章的结构体
//type Passage struct {
//	title,
//	time,
//	content,
//	url,
//	source string
//}

func main() {
	//sqlCommand := "INSERT INTO sea_news_todayhot_v2 SET news_title=?, news_date=?, news_content=?, news_web_url=?, news_source=?"
	//爬虫部分
	passages := tools.Crawler()//爬取新闻
	tools.ExecMysql(passages)//写入mysql

	//fmt.Println(passages)
}

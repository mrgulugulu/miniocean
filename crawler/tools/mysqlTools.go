package tools

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	_ "github.com/go-sql-driver/mysql" //注册这个驱动
)

const (
	driverName = "mysql"
	user = "root"
	passWord = "123456"
	dataBase = "test"
)

type News struct {
	Id uint `gorm:"AUTO_INCREMENT"`
	NewsTitle string
	NewsDate string
	NewsContent string
	NewsUrl string
	NewsSource string
}
//type Passage struct {
//	title,
//	time,
//	content,
//	url,
//	source string
//}

func ExecMysql(passages []Passage)  {

	//sqlCommand := "INSERT INTO sea_news_todayhot_v2 SET news_title=?, news_date=?, news_content=?, news_web_url=?, news_source=?"
	//进行数据库写入，使用go-sql-driver
	db, err := gorm.Open(driverName, fmt.Sprintf("%s:%s@/%s?charset=utf8", user, passWord, dataBase))
	//db, err := sql.Open(driverName, fmt.Sprintf("%s:%s@/%s?charset=utf8", user, passWord, dataBase))
	fmt.Println("Connecting mysql right now!")
	db.AutoMigrate(&News{})
	//db.AutoMigrate(&sea_news_todayhot_v2{})
	check(err)
	defer db.Close()
	for i := range passages {
		if checkIdentical(db, passages[i].Title) {
			//if err := db.Create(&sea_news_todayhot_v2{new: passages[i].Title, Time: passages[i].Time, Content: passages[i].Content, Url: passages[i].Url, Source: passages[i].Source}).Error; err != nil{
			//	log.Println("fail to create ")
			//} else {
			//	log.Println("create successfully")
			//}
			//stmt, err := db.Prepare(sqlCommand)
			//check(err)
			//res, err := stmt.Exec(passages[i].Title, passages[i].Time, passages[i].Content, passages[i].Url, passages[i].Source)
			//check(err)
			//_, err = res.LastInsertId()
			//check(err)
			//log.Printf("Title: %s news is written successfully!\n", passages[i].Title)
			info := News{NewsTitle: passages[i].Title, NewsDate: passages[i].Time, NewsContent: passages[i].Content,
				NewsUrl: passages[i].Url, NewsSource: passages[i].Source}
			if err := db.Create(&info).Error; err != nil {//如果有没错误的话就会返回nil，有错误就会返回error
				log.Printf("title: %s fail to create", passages[i].Title)
			} else {
				log.Printf("title: %s is created successfully!", passages[i].Title)
			}
		} else {
			log.Printf("Title: %s exists!\n", passages[i].Title)
			continue
		}
	}
	log.Println("Finished")
}
func checkIdentical(db *gorm.DB, content string) bool {
	var passage News
	if err := db.First(&passage ,"news_title=?", content).Error; err != nil {
		return true//找到就返回nil, 找不到就返回error
	}
	return false
}
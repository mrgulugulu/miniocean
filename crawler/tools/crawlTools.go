package tools

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

var (
	reOcean = `http://ocean.china.com.cn/[\d-\d]+[/[\w-./?%&=]*]?`
	reTitle = `<h1 class="artTitle">(.+)`
	reTime = `<div class="pub_date">发布时间：(.*?)</div>`
	reContent = `<p style="text-indent: 2em; margin-bottom: 15px;">(.*?)</p>`
	reSource = `<a href="http://ocean.china.com.cn/index.htm" target="_blank" class="" >(.*?)</a>&nbsp;>&nbsp;`
)

type Passage struct {
	Title,
	Time,
	Content,
	Url,
	Source string
}

func Crawler() []Passage {//将其导出
	wg := &sync.WaitGroup{}//同步池子
	finalPassage := []Passage{}
	resp, err := http.Get("http://ocean.china.com.cn/node_7198590.htm")
	check(err)
	defer resp.Body.Close()
	pageBytes, err := ioutil.ReadAll(resp.Body)
	check(err)
	//fmt.Println(string(pageBytes))
	pageStr := string(pageBytes)
	re := regexp.MustCompile(reOcean)
	urlResults := re.FindAllStringSubmatch(pageStr, -1)
	//fmt.Println(results)
	//t := time.Now()
	for _, url := range urlResults{
		wg.Add(1)
		eachPassage := new(Passage)
		go func(url []string, eachPassage Passage) {
			eachContent, eachTitle, eachTime, eachSource := extractContent(url)
			if len(eachContent) != 0 {
				eachContent = strings.Replace(eachContent, "</strong>", "", -1)
				eachPassage.Content = strings.Replace(eachContent, "<strong>", "", -1)
				eachPassage.Title = strings.Replace(eachTitle, "\r", "", -1)
				eachPassage.Time = eachTime
				eachPassage.Url = url[0]
				eachPassage.Source = eachSource
				finalPassage = append(finalPassage, eachPassage)
			}
			wg.Done()
		}(url, *eachPassage)
		//fmt.Println(finalContent)
	}
	wg.Wait()
	return finalPassage
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func extractContent(url []string) (string, string, string, string) {
	eachContent := ""
	subresp, _ := http.Get(url[0])//获取子url
	rawContents, err := ioutil.ReadAll(subresp.Body)
	check(err)
	contentsStr := string(rawContents)
	eachTitle := reResult(reTitle, contentsStr)[0][1]
	eachTime := reResult(reTime, contentsStr)[0][1]
	eachSource := reResult(reSource, contentsStr)[0][1]
	resultsOfContent := reResult(reContent, contentsStr)
	for i := range resultsOfContent {
		eachContent += resultsOfContent[i][1]
	}
	return eachContent, eachTitle, eachTime, eachSource
}

func reResult(re, content string) [][]string {
	reg := regexp.MustCompile(re)
	result := reg.FindAllStringSubmatch(content, -1)
	return result
}
package parserSJ

import (
	"awesomeProject/crawler/shijiejiayuan/type"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

//url示例shttp://www.jiayuan.com/59695744
const userUrlRe = `&p=([^&]+)&`

//用户列表页面请求内容解析
func ParserUserList(content, url string,cookies []*http.Cookie) types.ParserResult {
	re := regexp.MustCompile(userUrlRe)
	content = strings.Replace(content, "\\", "", -1)
	//fmt.Println("正文信息:", content)
	matchs := re.FindSubmatch([]byte(url))
	count, err := strconv.Atoi(string(matchs[1]))
	count++
	fmt.Printf("%d" ,count)
	//解析页面参数
	fmt.Printf(content)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		log.Printf("Parser User Error : %v", err)
	}

	doc.Find(" #normal_user_container > li > div > div.search_userHead > a.openBox.os_stat").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("jiexi url :%s ",s.Text())
	})

	result := types.ParserResult{}
	//for _, m := range matchs {
	//	result.Request = append(result.Request, models.Request{
	//		Url:        "http://" + string(m[1]),
	//		Cookies:cookies,
	//		Parserfunc: ParserUser,
	//	})
	//}
	return result
}

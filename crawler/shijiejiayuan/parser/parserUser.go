package parserSJ

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/model"
	"awesomeProject/crawler/shijiejiayuan/type"
	"awesomeProject/crawler_distributed/config"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
	"strings"
)

//用户解析
func ParserUser(resultUser, urls string) types.ParserResult {
	//fmt.Println("正文信息:", contents)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resultUser))
	if err != nil {
		log.Printf("Parser User Error : %v", err)
	}

	title := doc.Find(".member_name").Text()
	titles := strings.Split(title, "，")
	h4Test := doc.Find(".member_info_r h4")
	h4Tests := strings.Split(h4Test.Text(), ":")
	spanText := h4Test.Find("span").Text()

	var profile model.Profile
	profile.Name = strings.Replace(h4Tests[0], "ID", "", -1)
	arg, err := strconv.Atoi(strings.Replace(titles[0], "岁", "", -1))
	if err != nil {
		log.Printf("ages error : %v", err)
	}
	profile.Age = arg
	profile.Marriage = titles[1]
	profile.Hokou = strings.Replace(titles[2], "来自", "", -1)

	doc.Find(" div.member_info_r.yh > ul > li > div.fl.pr").Each(func(i int, s *goquery.Selection) {
		var str string
		if s.Find("em").Text() != "" {
			str = s.Find("em").Text()
		} else if s.Find("a").Text() != "" {
			str = s.Find("a").Text()
		}
		switch i {
		case 0: //学历
			profile.Occupation = str
		case 1: //身高
			str = strings.Replace(str, "cm", "", -1)
			height, err := strconv.Atoi(str)
			if err != nil {
				height = 0
			}
			profile.Height = height
		case 2: //购车
			profile.Car = str
		case 3: //月薪
			profile.Income = str
		case 4: //购房
			profile.House = str
		case 5: //体重
			str = strings.Replace(str, "斤", "", -1)
			weight, err := strconv.Atoi(str)
			if err != nil {
				weight = 0
			}
			profile.Weight = weight
		case 6: //星座
			profile.Xinzuo = str
		}
	})
	gender := doc.Find("body > div.subnav_box.yh > div > ul > li.cur > a").Text()
	if strings.Contains(gender, "她") {
		profile.Gender = "女士"
	} else {
		profile.Gender = "男士"
	}

	imageUrl, exists := doc.Find("#smallImg > div > ul > li > a > img").Attr("_src")
	if exists {
		profile.UserImageUrl = imageUrl
	} else {
		profile.UserImageUrl = "http://127.0.0.1:8888/blackUser.jpg"
	}

	result := types.ParserResult{
		Request: []types.Request{
			{
				Url:        config.GetUserUrl,
				Parserfunc: ParserJavascript,
			},
		},
		Items: []engine.Item{
			{
				Url:     urls,
				Type:    "shijijiayuan",
				Id:      strings.Replace(spanText, "ID:", "", -1),
				Payload: profile,
			},
		},
	}
	return result
}

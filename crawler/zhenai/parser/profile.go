package parser

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/model"
	"regexp"
	"strconv"
)

var marriageRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)</div>`)
var ageRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)岁</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)kg</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>月收入:([^<]+)</div>`) //收入
//var genderRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)cm</div>`)//性别
//var educationRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)</div>`)//职业
//var occupationRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)岁</div>`)//教育
var hokouRe = regexp.MustCompile(`<div class="m-btn pink" [^>]*>籍贯:([^>]+)</div>`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^>]+)\([^>]+\)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink" [^>]*>(已[^<]+房)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink" [^>]*>(已[^<]+车)</div>`)
var imagesUrlRe = regexp.MustCompile(`class="logo f-fl" style="background-image:url\(([^?|)]+)`)

var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func parserProfile(contents []byte, p ProfileParser) engine.ParserResult {
	var profile = model.Profile{} //用户信息结构体

	age, err := strconv.Atoi(extractString(contents, ageRe)) //年龄
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe)) //身高
	if err == nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe)) //体重
	if err == nil {
		profile.Weight = weight
	}

	profile.Marriage = extractString(contents, marriageRe) //婚况
	profile.Income = extractString(contents, incomeRe)     //收入
	profile.Gender = p.Gender                              //性别
	profile.Hokou = extractString(contents, hokouRe)       //户口
	profile.House = extractString(contents, houseRe)       //房
	profile.Car = extractString(contents, carRe)           //车
	//profile.Education = extractString(contents, educationRe) //职业
	profile.Xinzuo = extractString(contents, xinzuoRe) //星座
	//profile.Occupation = extractString(contents,occupationRe) //教育
	profile.Name = p.UserName //用户名

	profile.UserImageUrl = extractString(contents, imagesUrlRe)
	//log.Println(profile)
	result := engine.ParserResult{
		Items: []engine.Item{
			{
				Url:     p.Url,
				Type:    "zhenai",
				Id:      extractString([]byte(p.Url), idUrlRe),
				Payload: profile,
			},
		},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		//fmt.Printf("情况 : %s \n",match[1])
		return string(match[1])
	} else {
		return "-"
	}

}

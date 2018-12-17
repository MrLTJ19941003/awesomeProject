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
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>月收入:([^<]+)</div>`)//收入
//var genderRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)cm</div>`)//性别
//var educationRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)</div>`)//职业
//var occupationRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)岁</div>`)//教育
var hokouRe = regexp.MustCompile(`<div class="m-btn pink" [^>]*>籍贯:([^>]+)</div>`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^>]+)\([^>]+\)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink" [^>]*>(已[^<]+房)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink" [^>]*>(已[^<]+车)</div>`)

func ParserProfile(contents []byte) engine.ParserResult {
	var profile= model.Profile{} //用户信息结构体

	age, err := strconv.Atoi(extractString(contents,ageRe))  //年龄
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
	profile.Income = extractString(contents, incomeRe) //收入
	//profile.Gender = extractString(contents, genderRe) //性别
	profile.Hokou = extractString(contents, hokouRe) //户口
	profile.House = extractString(contents, houseRe) //房
	profile.Car = extractString(contents, carRe) //车
	//profile.Education = extractString(contents, educationRe) //职业
	profile.Xinzuo = extractString(contents, xinzuoRe) //星座
	//profile.Occupation = extractString(contents,occupationRe) //教育
	//profile.Name = extractString(contents, marriageRe) //用户名

	//log.Println(profile)
	return engine.ParserResult{
		Items:[]interface{}{profile},
	}
}

func extractString(contents []byte,re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match)  >= 2 {
		//fmt.Printf("情况 : %s \n",match[1])
		return string(match[1])
	}else{
		return ""
	}

}

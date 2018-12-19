package parser

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler_distributed/config"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
	genderRe  = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^>]+)</td>`)
)

func ParserCity(contents []byte) engine.ParserResult {
	//re := regexp.MustCompile(cityRe)
	matches := profileRe.FindAllSubmatch(contents, -1)
	genderMatches := genderRe.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for i, m := range matches {
		//result.Items = append(result.Items,"User "+string(m[2]))
		/*url := string(m[1])
		name := string(m[2])
		gender :=  string(genderMatches[i][1])*/
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: NewProfileParser(string(m[2]), string(genderMatches[i][1]), string(m[1])),
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParserCity, config.ParserCity),
		})
	}

	return result
}

//func ProfileParser(name string,gender string,url string) engine.ParserFunc{
//	return func(b []byte) engine.ParserResult{
//		return ParserProfile(b,name,gender,url)
//	}
//}

type ProfileParser struct {
	UserName string
	Gender   string
	Url      string
}

func NewProfileParser(userName, gender, url string) ProfileParser {
	return ProfileParser{
		UserName: userName,
		Gender:   gender,
		Url:      url,
	}
}

func (p ProfileParser) Parser(contents []byte) engine.ParserResult {
	return parserProfile(contents, p)
}

func (p ProfileParser) Serialize() (name string, args []interface{}) {
	//m := make(map[string]string)
	//m["UserName"] = p.UserName
	//m["Gender"] = p.Gender
	//m["Url"] = p.Url
	var s []interface{}
	s = append(s, p.UserName)
	s = append(s, p.Gender)
	s = append(s, p.Url)
	return config.ParserProfile, s
}

//func FromJsonObj(o interface{})(ProfileParser ,error){
//	var profile ProfileParser
//	s, err := json.Marshal(o)
//	if err != nil {
//		return profile,err
//	}
//	err = json.Unmarshal(s,&profile)
//	return profile,err
//}

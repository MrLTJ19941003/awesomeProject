package parser

import (
	"awesomeProject/crawler/engine"
	"regexp"
)

var (
	profileRe  = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
	genderRe = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^>]+)</td>`)
)
func ParserCity(contents []byte)engine.ParserResult{
	//re := regexp.MustCompile(cityRe)
	matches := profileRe.FindAllSubmatch(contents, -1)
	genderMatches := genderRe.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for i,m := range matches{
		//result.Items = append(result.Items,"User "+string(m[2]))
		/*url := string(m[1])
		name := string(m[2])
		gender :=  string(genderMatches[i][1])*/
		result.Requests = append(result.Requests,engine.Request{
			Url : string(m[1]),
			ParserFunc : ProfileParser(string(m[2]),string(genderMatches[i][1]),string(m[1])),
		})
	}

	matches =cityUrlRe.FindAllSubmatch(contents, -1)
	for _,m := range matches{
		result.Requests = append(result.Requests,engine.Request{
			Url : string(m[1]),
			ParserFunc : ParserCity,
		})
	}

	return result
}

func ProfileParser(name string,gender string,url string) engine.ParserFunc{
	return func(b []byte) engine.ParserResult{
		return ParserProfile(b,name,gender,url)
	}
}

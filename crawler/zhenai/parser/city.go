package parser

import (
	"awesomeProject/crawler/engine"
	"regexp"
)

var (
	profileRe  = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)
func ParserCity(contents []byte)engine.ParserResult{
	//re := regexp.MustCompile(cityRe)
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _,m := range matches{
		//result.Items = append(result.Items,"User "+string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url : string(m[1]),
			ParserFunc : ParserProfile,
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

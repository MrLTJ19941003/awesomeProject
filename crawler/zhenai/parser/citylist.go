package parser

import (
	"awesomeProject/crawler/engine"
	"regexp"
)

const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)"[^>]*>([^<]+)</a>`
func ParseCityList(contents []byte) engine.ParserResult{
	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	//count :=0
	for _,m := range matchs{
		//result.Items = append(result.Items,"City "+string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url : string(m[1]),
			ParserFunc : ParserCity,
		})
		/*count++
		if count>10{
			break
		}*/
	}
	return result
}



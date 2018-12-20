package parserSJ

import (
	"awesomeProject/crawler/shijiejiayuan/type"
	"regexp"
	"strings"
)

//url示例shttp://www.jiayuan.com/59695744
const userUrlRe = `href="http://(www.jiayuan.com/[0-9]+)?`

//ajax接口请求内容解析
func ParserJavascript(content, url string) types.ParserResult {
	re := regexp.MustCompile(userUrlRe)
	content = strings.Replace(content, "\\", "", -1)
	//fmt.Println("正文信息:", content)
	matchs := re.FindAllSubmatch([]byte(content), -1)

	result := types.ParserResult{}
	for _, m := range matchs {
		result.Request = append(result.Request, types.Request{
			Url:        "http://" + string(m[1]),
			Parserfunc: ParserUser,
		})
	}
	return result
}

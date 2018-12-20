package main

import (
	"awesomeProject/crawler/shijiejiayuan/engine"
	"awesomeProject/crawler/shijiejiayuan/fetcher"
	"awesomeProject/crawler/shijiejiayuan/parser"
	"awesomeProject/crawler/shijiejiayuan/scheduler"
	"awesomeProject/crawler/shijiejiayuan/type"
	"awesomeProject/crawler_distributed/config"
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"gopkg.in/olivere/elastic.v5"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

//世纪佳缘网数据（单机版爬虫启动入口）
func main() {
	client := &http.Client{}
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client.Jar = jar //设置cookie
	//login(client)
	elasticClient, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	//simpleEngine := engineSJ.SimpleEngine{ //单机版爬虫
	//	ClientElastic: elasticClient,
	//	ClientHttp:    client,
	//}

	concurrentEngine := engineSJ.ConcurrentEngine{ //并发简单版爬虫
		WorkerCount:   100,
		ClientElastic: elasticClient,
		ClientHttp:    client,
		Scheduler:     	&schedulerSJ.SimpleSchedulerSJ{},
	}

	requestUserContext := types.Request{
		Url:        "http://www.jiayuan.com/187086135?fxly=search_v2_index",
		Parserfunc: parserSJ.ParserUser,
	}
	requestGetUserUrl := types.Request{
		Url:        config.GetUserUrl,
		Parserfunc: parserSJ.ParserJavascript,
	}
	concurrentEngine.Run(requestUserContext, requestGetUserUrl)
	//simpleEngine.Run(requestUserContext, requestGetUserUrl)
}

//登陆
func login(client *http.Client) {
	url := "https://passport.jiayuan.com/dologin.php?pre_url=http://www.jiayuan.com/usercp/"
	fromparams := make(map[string]interface{})
	fromparams["name"] = "19928709269"
	fromparams["password"] = "liu01234"
	fromparams["remem_pass"] = ""
	//reuqest := types.CreateReuqest(method, url, fromparams)
	fetcher.LoginRequest(fromparams, client, url)
}

//gbk转换成utf-8
func gbk2utf8(str []byte) ([]byte, error) {
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewDecoder()))
}

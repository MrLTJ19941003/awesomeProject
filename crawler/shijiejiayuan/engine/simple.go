package engineSJ

import (
	"awesomeProject/crawler/shijiejiayuan/elasticdb"
	"awesomeProject/crawler/shijiejiayuan/fetcher"
	"awesomeProject/crawler/shijiejiayuan/type"
	"awesomeProject/crawler_distributed/config"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"net/http"
)

type SimpleEngine struct {
	ClientElastic *elastic.Client
	ClientHttp    *http.Client
}

func (s SimpleEngine) Run(seeds ...types.Request) {
	var requests []types.Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		parserResultUser, err := fetcher.SendRequests(request, s.ClientHttp)
		if err != nil {
			log.Printf("error : %v", err)
			return
		}

		parseresult := request.Parserfunc(parserResultUser, request.Url)
		for _, item := range parseresult.Items {
			log.Printf("item : %v ", item)
			err = elasticdb.Saver(item, s.ClientElastic, config.ElasticIndex)
			if err != nil {
				log.Printf("saver error item : %v: %v", item, err)
			}
		}

		for _, req := range parseresult.Request {
			//去除重复数据
			duplicate := isDuplicate(req.Url)
			if !duplicate {
				log.Printf("request Url : %v ", req)
				requests = append(requests, req)
			}
		}
	}
}

//去除用户重复数据
var visitedUrl = make(map[string]bool)

func isDuplicate(url string) bool {
	if url == "http://www.jiayuan.com/ajax/interested.php?r=0.32553824808613996&ad_param[]=pid%3Apersonalmatch_profile_new%7Ccount%3A6%7C"{
		return false
	}
	if visitedUrl[url] {
		return true
	}
	visitedUrl[url] = true
	return false
}

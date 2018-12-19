package main

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/zhenai/parser"
	"awesomeProject/crawler_distributed/config"
	"awesomeProject/crawler_distributed/rpcsupport"
	"awesomeProject/crawler_distributed/worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServerRpc(host, &worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	request := engine.Request{
		Url:    "http://album.zhenai.com/u/1690271375",
		Parser: parser.NewProfileParser("安静的雪", "女士", "http://album.zhenai.com/u/1690271375"),
	}
	var result worker.ParserResult
	err = client.Call(config.CrawRpc, worker.SerializeRequeset(request), &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}

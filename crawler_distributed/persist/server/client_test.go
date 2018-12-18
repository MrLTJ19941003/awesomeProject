package main

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/model"
	"awesomeProject/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T){
	//start ItemSaverServer
	host := ":1234"
	go serveRpc(host,"test1")
	time.Sleep(time.Second)
	//start ItemSaverClient
	client,err := rpcsupport.NewClient(host)
	if err != nil {
		//log.Printf("client error : %v ",err)
		panic(err)
	}
	//call save
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/93861579",
		Type: "zhenai",
		Id:   "93861579",
		Payload: model.Profile{
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3000-5000元",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
			Marriage:   "离异",
			House:      "已购房",
			Hokou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}
	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)
	if err!=nil || result != "ok"{
		t.Errorf("result : %s ;error : %s",result,err)
	}
}
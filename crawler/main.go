package main

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/persist"
	"awesomeProject/crawler/scheduler"
	"awesomeProject/crawler/zhenai/parser"
	"awesomeProject/crawler_distributed/config"
)

func main() {
	itemChan, err := persist.ItemSaver(config.ElasticIndex)
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		//Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	//e := engine.SimpleEngine{
	//}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParserCityList), //parser.ParseCityList,
		//Url : "http://www.zhenai.com/zhenghun/shanghai",
		//Parser : engine.NewFuncParser(parser.ParserCity,config.ParserCity),
		//Url : "http://album.zhenai.com/u/1690271375",
		//Parser :parser.NewProfileParser("安静的雪","女士","http://album.zhenai.com/u/1690271375"),
	})
}

package main

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/persist"
	"awesomeProject/crawler/scheduler"
	"awesomeProject/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		//Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:100,
		ItemChan: persist.ItemSaver(),
	}
	//e := engine.SimpleEngine{
	//}
	e.Run(engine.Request{
		//Url : "http://www.zhenai.com/zhenghun",
		//ParserFunc : parser.ParseCityList,
		Url : "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc : parser.ParserCity,
		//Url : "http://album.zhenai.com/u/1690271375",
		//ParserFunc : parser.ParserProfile,
	})
}

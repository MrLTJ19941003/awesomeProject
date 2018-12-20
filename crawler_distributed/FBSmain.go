package main

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/scheduler"
	"awesomeProject/crawler/zhenai/parser"
	"awesomeProject/crawler_distributed/config"
	"awesomeProject/crawler_distributed/persist/client"
	"awesomeProject/crawler_distributed/rpcsupport"
	"awesomeProject/crawler_distributed/worker/client"
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host","","itemsaver host")
	workerHosts = flag.String("worker_hosts","","worker hosts (comma separated)")
)

//分布式爬虫入口
func main() {
	flag.Parse()
	if *itemSaverHost == "" || *workerHosts == ""{
		fmt.Println("Must specify a itemSaverHost and workerHosts")
		return
	}
	//初始化persist存储RPC客户端
	itemChan, err := persistclient.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	client := createClientPool(strings.Split(*workerHosts,","))
	//初始化worker爬取RPC客户端
	processor := workerclient.CreateProcessor(client)

	//初始化爬虫engine引擎
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		//Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
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

func createClientPool(hosts []string) chan *rpc.Client{
	var clients []*rpc.Client
	for _,h := range hosts{
		client ,err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients,client)
			log.Printf("Connecting to %s",h)
		}else{
			log.Printf("Error connecting to %s : %v",h,err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for{
			for _,client := range clients{
				out <- client
			}
		}
	}()
	return out
}

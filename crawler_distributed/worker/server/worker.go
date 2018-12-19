package main

import (
	"awesomeProject/crawler_distributed/rpcsupport"
	"awesomeProject/crawler_distributed/worker"
	"flag"
	"fmt"
)

var port = flag.Int("port",0,"the port for me to me listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Must specify a port")
		return
	}
	err := rpcsupport.ServerRpc(fmt.Sprintf(":%d", *port), &worker.CrawlService{})
	if err != nil {
		panic(err)
	}
}

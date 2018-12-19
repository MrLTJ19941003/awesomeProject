package main

import (
	"awesomeProject/crawler_distributed/config"
	"awesomeProject/crawler_distributed/persist"
	"awesomeProject/crawler_distributed/rpcsupport"
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
)

var port = flag.Int("port",0,"the port for me to me listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Must specify a port")
		return
	}
	err := serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndexTest)
	if err != nil {
		panic(err)
	}

}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return rpcsupport.ServerRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}

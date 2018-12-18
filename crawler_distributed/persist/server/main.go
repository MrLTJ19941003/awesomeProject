package main

import (
	"awesomeProject/crawler_distributed/persist"
	"awesomeProject/crawler_distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
)

func main(){
	err := serveRpc(":8990", "dating_profile")
	if err != nil {
		panic(err)
	}

}

func serveRpc(host ,index string) error {
	client , err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil{
		panic(err)
	}

	return rpcsupport.ServerRpc(host,
		&persist.ItemSaverService{
			Client:client,
			Index:index,
		})
}
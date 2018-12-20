package elasticdb

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/persist"
	"gopkg.in/olivere/elastic.v5"
)

func Saver(item engine.Item, client *elastic.Client, index string) error {
	return persist.Save(item,client,index)
}


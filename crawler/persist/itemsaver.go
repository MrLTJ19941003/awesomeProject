package persist

import (
	"awesomeProject/crawler/engine"
	"context"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver : Got item #%d : %v ", itemCount, item)
			itemCount++

			Save(item, client, index)
		}
	}()
	return out, nil
}

func Save(item engine.Item, client *elastic.Client, index string) error {

	if item.Type == "" {
		return errors.New("must supply type!")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type)

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}

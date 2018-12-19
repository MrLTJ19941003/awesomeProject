package persist

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/persist"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, resule *string) error {
	err := persist.Save(item, s.Client, s.Index)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*resule = "ok"
	} else {
		log.Printf("error saving item %v :%v.", item, err)
	}
	return err
}

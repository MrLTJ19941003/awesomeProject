package persist

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/persist"
	"gopkg.in/olivere/elastic.v5"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(item engine.Item,resule *string) error {
	err := persist.Save(item,s.Client,s.Index)
	if err == nil {
		*resule = "ok"
	}
	return err
}

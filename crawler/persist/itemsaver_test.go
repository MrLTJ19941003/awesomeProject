package persist

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/model"
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestSave(t *testing.T) {
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/93861579",
		Type: "zhenai",
		Id:   "93861579",
		Payload: model.Profile{
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3000-5000元",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
			Marriage:   "离异",
			House:      "已购房",
			Hokou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}
	client , err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil{
		panic(err)
	}

	const index  = "dating_test"
	err = Save(item,client,index)
	if err != nil{
		panic(err)
	}

	resp, err := client.Get().Index(index).Type(item.Type).Id(item.Id).Do(context.Background())
	if err != nil{
		panic(err)
	}
	t.Logf("%s" ,*resp.Source)
	//b := []byte(*resp.Source)
	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil{
		panic(err)
	}
	actualProfile ,err := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != item {
		t.Errorf("got %v, expected %v",
			actual,item)
	}
}

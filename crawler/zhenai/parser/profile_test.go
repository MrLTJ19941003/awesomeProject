package parser

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/fetcher"
	"awesomeProject/crawler/model"
	"fmt"
	"testing"
)

func TestParserProfile(t *testing.T) {

	item := engine.Item{
		Url:  "http://album.zhenai.com/u/93861579",
		Type: "zhenai",
		Id:   "93861579",
		Payload: model.Profile{
			UserImageUrl: "https://photo.zastatic.com/images/photo/23466/93861579/16631621612264856.jpg",
			Age:          34,
			Height:       162,
			Weight:       57,
			Income:       "3000-5000元",
			Gender:       "女",
			Name:         "安静的雪",
			Xinzuo:       "牡羊座",
			Occupation:   "人事/行政",
			Marriage:     "离异",
			House:        "已购房",
			Hokou:        "山东菏泽",
			Education:    "大学本科",
			Car:          "未购车",
		},
	}
	contents, err := fetcher.Fetch("http://album.zhenai.com/u/1670025259")
	if err != nil {
		panic(err)
	}
	parserResult := parserProfile(contents, ProfileParser{
		UserName: "安静的雪",
		Url:      "http://album.zhenai.com/u/93861579",
		Gender:   "女",
	})

	fmt.Printf("result : %v ,\n item : %v", parserResult, item)
}

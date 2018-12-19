package view

import (
	"awesomeProject/crawler/engine"
	searchModel "awesomeProject/crawler/frontend/model"
	"awesomeProject/crawler/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) { //SearchResultView测试方式
	view := CreateSearchResultView("search.html") //将html页面加载至template模块中

	out, err := os.Create("template.test.html") //创建一个html页面
	page := searchModel.SearchResult{}
	page.Hits = 120
	page.Start = 70
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
	for i := 0; i <= 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page) //将page结构体对象数据映射进html页面中，然后再加载至输出流中。
	if err != nil {
		panic(err)
	}
}

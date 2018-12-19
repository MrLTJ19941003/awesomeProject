package controller

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/frontend/model"
	"awesomeProject/crawler/frontend/view"
	"context"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	//SearchResultHandler结构体
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler { //初始化SearchResultHandler
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

// locahost:8888/search?q=男 已购房&from20
func (s SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) { //实现handler接口需要实现的方法
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}

	var page model.SearchResult
	page, err = s.getSearchResult(q, from)
	if err != nil { //出现错误，将异常抛出给http服务器展示给用户
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = s.view.Render(w, page)

	if err != nil { //出现错误，将异常抛出给http服务器展示给用户
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/**
传入需要查询的参数，调用ElasticSearch查询接口，然后将查询出来的结果组装至page结构体中
 */
func (s SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := s.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).
		Do(context.Background())
	if err != nil {
		return result, err
	}
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Querystr = q

	//两种方式将 resp.Each(reflect.TypeOf(engine.Item{})) 放入 result.Items中
	//第一种是用range遍历，然后通过v.(type)类型转换append进行Items中
	//第二种是将items的类型改为interface{}类型。这里选用第二种方式
	//for _,v := range resp.Each(reflect.TypeOf(engine.Item{})){
	//	result.Items  = append(result.Items,v.(engine.Item))
	//}
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)
	return result, nil
}

func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}

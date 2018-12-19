package model

type SearchResult struct {
	//搜索结果结构体
	Hits     int64
	Start    int
	Querystr string
	PrevFrom int
	NextFrom int
	Items    []interface{} //[]engine.Item
}

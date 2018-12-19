package engine

import "awesomeProject/crawler_distributed/config"

type ParserFunc func([]byte) ParserResult

type Parser interface {
	Parser(contents []byte) ParserResult
	Serialize() (name string, args []interface{})
}

type Request struct {
	Url    string
	Parser Parser
}

type ParserResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type NilParser struct{}

func (NilParser) Parser(_ []byte) ParserResult {
	return ParserResult{}
}

func (NilParser) Serialize() (name string, args []interface{}) {
	return config.NilParser, nil
}

//func NilParser([]byte) ParserResult {
//	return ParserResult{}
//}

type FuncParser struct {
	parser ParserFunc
	Name   string
}

func (f *FuncParser) Parser(contents []byte) ParserResult {
	return f.parser(contents)
}

func (f *FuncParser) Serialize() (name string, args []interface{}) {
	return f.Name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		Name:   name,
	}
}

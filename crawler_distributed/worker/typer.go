package worker

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/zhenai/parser"
	"awesomeProject/crawler_distributed/config"
	"errors"
	"fmt"
	"log"
)

type SerializedParser struct {
	Name string
	Args []interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParserResult struct {
	Item    []engine.Item
	Request []Request
}

func SerializeRequeset(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParserResult) ParserResult {
	result := ParserResult{
		Item: r.Items,
	}
	for _, req := range r.Requests {
		result.Request = append(result.Request, SerializeRequeset(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := DeserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(r ParserResult) engine.ParserResult {
	result := engine.ParserResult{
		Items: r.Item,
	}
	for _, req := range r.Request {
		engineRequest, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request : %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineRequest)
	}
	return result
}

func DeserializeParser(s SerializedParser) (engine.Parser, error) {
	switch s.Name {
	case config.ParserCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParserCityList), nil
	case config.ParserCity:
		return engine.NewFuncParser(parser.ParserCity, config.ParserCity), nil
	case config.ParserProfile:
		var profileParsers []string
		for _, ss := range s.Args {
			if ss, ok := ss.(string); ok {
				profileParsers = append(profileParsers, ss)
			} else {
				//parfileParsers = append(parfileParsers, "")
				return nil, fmt.Errorf("error : invalid arg : %v ", s.Args)
			}
		}
		return parser.NewProfileParser(profileParsers[0], profileParsers[1], profileParsers[2]), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unknown parser name")
	}
}

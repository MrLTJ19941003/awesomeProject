package types

import "awesomeProject/crawler/engine"

type SchedulerSJ interface {
	engine.ReadyNotifier
	Submits(Request)
	WorkerChan() chan Request
	Runs()
}

type ParserFunc func(content ,url string) ParserResult

type Request struct {
	Url string
	Parserfunc ParserFunc
}

type ParserResult struct {
	Request []Request
	Items []engine.Item
}
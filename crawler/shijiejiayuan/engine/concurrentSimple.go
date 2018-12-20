package engineSJ

import (
	"awesomeProject/crawler/shijiejiayuan/fetcher"
	"awesomeProject/crawler/shijiejiayuan/type"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"net/http"
)

type ConcurrentEngine struct {
	ClientElastic *elastic.Client
	ClientHttp    *http.Client
	WorkerCount   int
	Scheduler     types.SchedulerSJ
}

func (c ConcurrentEngine) Run(seeds ...types.Request) {
	c.Scheduler.Runs()


	out := make(chan types.ParserResult)

	for i := 0; i < c.WorkerCount; i++ {
		c.createWorker(c.Scheduler.WorkerChan(), out)
	}

	for _,req := range seeds{
		c.Scheduler.Submits(req)
	}

	for {
		parseresult := <- out
		for _, item := range parseresult.Items {
			log.Printf("item : %v ", item)
			//err := elasticdb.Saver(item, c.ClientElastic, config.ElasticIndex)
			//if err != nil {
			//	log.Printf("saver error item : %v: %v", item, err)
			//}
		}
		for _, req := range parseresult.Request {
			//去除重复数据
			duplicate := isDuplicate(req.Url)
			//fmt.Println( " ===========   ",req.Url , duplicate )
			if !duplicate {
				log.Printf("request Url : %v ", req)
				c.Scheduler.Submits(req)
			}
		}
	}
}

func (c ConcurrentEngine) createWorker(in chan types.Request,out chan types.ParserResult){
	go func() {
		for{
			request := <- in
			parserResultUser, err := fetcher.SendRequests(request, c.ClientHttp)
			if err != nil {
				log.Printf("fetcher error : %v" , err)
				continue
			}
			result := request.Parserfunc(parserResultUser, request.Url)
			out <- result
		}
	}()

}

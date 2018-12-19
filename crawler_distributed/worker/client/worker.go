package workerclient

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler_distributed/config"
	"awesomeProject/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(client chan *rpc.Client) engine.Processor{
	//client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.ClawPort0))

	return func(r engine.Request) (result engine.ParserResult, e error) {

		sreq := worker.SerializeRequeset(r)

		var results worker.ParserResult
		c := <- client
		err := c.Call(config.CrawRpc, sreq, &results)
		if err != nil {
			return engine.ParserResult{},err
		}
		result = worker.DeserializeResult(results)
		return result,nil
	}
}

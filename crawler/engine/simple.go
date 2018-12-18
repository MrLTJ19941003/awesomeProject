package engine

import (
	"log"
)

type SimpleEngine struct {

}

func (s SimpleEngine) Run(seeds ...Request){

	var requests []Request
	for _,r := range seeds{
		requests = append(requests,r)
	}

	for len(requests) >0{
		r := requests[0]
		requests = requests[1:]

		parserResults,err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests,
						parserResults.Requests...)
		for _,item := range parserResults.Items{
			log.Printf("Got item %v ",item)
		}
	}
}



package scheduler

import (
	"awesomeProject/lang/fastSearchFiles/models"
)

type CourrentQueuedScheduler struct {
	requestChan chan typesF.Workers
	workerChan  chan chan typesF.Workers
}

func (c *CourrentQueuedScheduler) Submit(workers typesF.Workers) {
	c.requestChan <- workers
}

func (c *CourrentQueuedScheduler) WorkerChan() chan typesF.Workers {
	return make(chan typesF.Workers)
}

func (c *CourrentQueuedScheduler) ReadyWorker(workers chan typesF.Workers) {
	c.workerChan <- workers
}

func (c *CourrentQueuedScheduler) Run() {
	c.workerChan = make(chan chan typesF.Workers)
	c.requestChan = make(chan typesF.Workers)
	go func() {
		var workerChanQ []chan typesF.Workers
		var requestChanQ []typesF.Workers
		for{
			var activeworkerChanQ chan typesF.Workers
			var activerequestChanQ typesF.Workers
			if len(workerChanQ) > 0 && len(requestChanQ) > 0{
				activeworkerChanQ = workerChanQ[0]
				activerequestChanQ = requestChanQ[0]
			}
			select {
			case r:=<-c.workerChan:
				workerChanQ = append(workerChanQ, r)
			case r:=<-c.requestChan:
				requestChanQ = append(requestChanQ, r)
			case activeworkerChanQ <- activerequestChanQ:
				workerChanQ = workerChanQ[1:]
				requestChanQ = requestChanQ[1:]
			}
		}
	}()
}





package schedulerSJ

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/shijiejiayuan/type"
)

type SimpleSchedulerSJ struct {
	workerChan chan types.Request
}

func (s *SimpleSchedulerSJ) WorkerReady(chan engine.Request) {
	panic("implement me")
}

func (s *SimpleSchedulerSJ) Submits(r types.Request) {
	s.workerChan <- r
}

func (s *SimpleSchedulerSJ) WorkerChan() chan types.Request {
	return s.workerChan
}

func (s *SimpleSchedulerSJ) Runs() {
	s.workerChan = make(chan types.Request )
}








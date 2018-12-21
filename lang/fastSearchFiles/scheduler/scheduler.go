package scheduler

import (
	"awesomeProject/lang/fastSearchFiles/models"
)

type Scheduler interface {
	Submit(workers typesF.Workers)
	Run()
	WorkerChan() chan typesF.Workers
	ReadyNotifierF
}

type ReadyNotifierF interface {
	ReadyWorker(workers chan typesF.Workers)
}

type StopChan interface {
	Stop() chan string
}


type SimpleScheduler struct {
	workerChan chan typesF.Workers
}


func (s *SimpleScheduler) ReadyWorker(workers chan typesF.Workers) {

}

func (s *SimpleScheduler) WorkerChan() chan typesF.Workers {
	return s.workerChan
}

func (s *SimpleScheduler) Submit(workers typesF.Workers) {
	go func() {
		s.workerChan <- workers
	}()
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan typesF.Workers)
}



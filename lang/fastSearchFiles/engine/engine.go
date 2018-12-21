package engine

import (
	"awesomeProject/lang/fastSearchFiles/models"
	"awesomeProject/lang/fastSearchFiles/scheduler"
	"log"
	"sync"
	"time"
)

type SimpleEngine struct {
	Scheduler scheduler.Scheduler
	WorkerCount int
}

func (s *SimpleEngine) Run(outfile chan []string ,seeds ...typesF.Workers){
	var wg sync.WaitGroup
	s.Scheduler.Run()
	out := make(chan typesF.WorkerResult)
	for i:=0;i<s.WorkerCount;i++{
		wg.Add(1)
		s.CreateWorker(s.Scheduler.WorkerChan(),out,s.Scheduler)
		//wg.Done()
	}

	for _,w := range seeds{
		s.Scheduler.Submit(w)
	}
	count := 1
	var str []string
	for{
		select {
		case result :=<- out:
			for _,f := range result.FindsName{
				//fmt.Printf("find file  %d : %v \n" ,count,f)
				str = append(str, f)
				count++
			}
			for _,w := range result.FilesName{
				s.Scheduler.Submit(w)
			}
		case <- time.After(time.Second)://goroutine超时设置
			outfile <-str
			return
		}
	}
	wg.Wait()
}

func  (s *SimpleEngine) CreateWorker(in chan typesF.Workers,out chan typesF.WorkerResult,read scheduler.ReadyNotifierF) {
	go func() {
		for{
			read.ReadyWorker(in)
			select {
			case result :=<- in:
				worderResult,err := result.WorderFunc(result)
				if err != nil{
					log.Printf("worker error : %v ",err)
				}
				out <- worderResult
			case <- time.After(time.Second):
				return
				//goto Loop
			}
		}
		//Loop:
		//	fmt.Println("退出for死循环")
	}()
}

package main

import (
	"awesomeProject/lang/fastSearchFiles/engine"
	"awesomeProject/lang/fastSearchFiles/models"
	scheduler2 "awesomeProject/lang/fastSearchFiles/scheduler"
	"awesomeProject/lang/fastSearchFiles/worker"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	outfile := make(chan []string)
	go func() {
		wg.Add(1)
		ss := <- outfile
		for _,s := range ss{
			fmt.Printf("%s \n",s)
		}
		wg.Done()
	}()

	scheduler := &scheduler2.SimpleScheduler{}
	//scheduler := &scheduler2.CourrentQueuedScheduler{}
	simpleEngine := engine.SimpleEngine{
		Scheduler:   scheduler,
		WorkerCount: 100,
	}
	simpleEngine.Run(outfile,typesF.Workers{
		Fliepath:   "f:/",
		WorderFunc: worker.FindFiles,
	})
	wg.Wait()


}

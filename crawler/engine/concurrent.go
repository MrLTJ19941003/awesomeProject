package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func  (c *ConcurrentEngine) Run(seeds ...Request)  {
	out := make(chan ParserResult)
	c.Scheduler.Run()

	for i :=0;i<c.WorkerCount;i++{
		createWorker(c.Scheduler.WorkerChan(), out,c.Scheduler)
	}

	for _,r := range seeds{
		if isDuplicate(r.Url){
			continue
		}
		c.Scheduler.Submit(r)
	}

	for {
		result := <- out
		for _,item := range result.Items{
			go func() {c.ItemChan <- item}()
		}

		for _,request := range result.Requests{
			if isDuplicate(request.Url){
				continue
			}
			c.Scheduler.Submit(request)
		}
	}
}

var visitedUrl = make(map[string]bool)
func isDuplicate(url string) bool{
	if visitedUrl[url] {
		return true
	}
	visitedUrl[url] = true
	return false
}

func createWorker(in chan Request,out chan ParserResult,ready ReadyNotifier){
	go func() {
		for {
			ready.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
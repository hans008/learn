package engine

import "log"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}


func (e *ConcurrentEngine) Run(seeds ...Request){
	//scheduler创建多个worker去处理任务

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i :=0;i<e.WorkerCount;i++{
		createWorker(out,e.Scheduler)
	}


	//将request扔给scheduler处理
	for _,r := range seeds {
		e.Scheduler.Submit(r)
	}

	//循环channel 将channel中的request交给scheduler去调度

	countItem :=0
	result := <- out

	//循环处理channel
	for {

		//对item计数
		for _,item := range result.Items{
			log.Printf("Got item #%d  %v \n",countItem,item)
			countItem ++
		}


		//将request交给scheduler继续调度
		for _, request := range result.Requests {
			log.Printf("Get Url %s",request.Url)
			e.Scheduler.Submit(request)
		}
	}

}


func createWorker(out chan ParseResult,s Scheduler){
	//每一个worker都是一个goroutine,worker功能是fetcher和parser

	in := make(chan Request)

	go func(){
		for {
			//tell scheduler i am ready
			s.WorkerReady(in)

			//从in 这个channel中获取request, fetcher处理后，将结果放回out 这个channel中
			request := <- in
			result,err :=worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

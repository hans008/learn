package scheduler

import "hans/learn/spider/crawler/engine"

type QueuedScheduler struct {
	requestChan chan  engine.Request
	//为每一个worker创建一个channel chan engine.Request
	//把所有的worker对应的channel 放入一个总的chanel,chan chan engine.Request
	workerChan chan chan engine.Request
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request){
	//有一个worker channel准备好了，可以去接收request请求了
	s.workerChan <- w
}

func (s *QueuedScheduler) ConfigMasterWorkerChan(chan engine.Request) {
	panic("implement me")
}

func (s *QueuedScheduler) Run(){
	//总控制的goroutine

	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func(){
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(activeWorker) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select {
				case r := <- s.requestChan:
					requestQ = append(requestQ,r)
				case w := <- s.workerChan:
					workerQ = append(workerQ,w)
				case activeWorker <- activeRequest:
					workerQ = workerQ[1:]
					requestQ = requestQ[1:]
			}
		}
	}()
}


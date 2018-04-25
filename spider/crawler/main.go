package main

import (
	"hans/learn/spider/crawler/engine"
	"hans/learn/spider/crawler/zhenai/parser"
	"hans/learn/spider/crawler/scheduler"
)

func main() {
	e :=engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount:10,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
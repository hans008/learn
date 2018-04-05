package engine

import (
	"hans/learn/spider/crawler/fetcher"
	"log"
)


type SimpleEngine struct {}


func (e SimpleEngine) Run(seeds ...Request){

	var requests  []Request

	for _,r := range seeds {
		requests = append(requests,r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests =requests[1:]

		log.Printf("Fetching %s",r.Url)

		parseResult,err := worker(r)
		if err != nil {
			continue
		}

		requests = append(requests,parseResult.Requests...)

		for _,item := range parseResult.Items {
			log.Printf("Got item %v",item)
		}



	}
}

func  worker(request Request) (ParseResult,error){
		body,err := fetcher.Fetch(request.Url)
		if err != nil {
			log.Printf("Fetcher: error " +
				"fetching url %s: %v",request.Url,err)
			return ParseResult{},err
		}

		return  request.ParserFunc(body),nil

}
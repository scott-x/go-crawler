package engine

import (
	"log"

	"github.com/clawler/fetcher"
)

func Run(seed ...Request){
	var requests []Request
	for _, r := range seed{
		requests = append(requests,r)
	}
	for len(requests)>0{
		r := requests[0]
		requests = requests[1:]
		log.Printf("fetching url %s",r.Url)
		body,err :=fetcher.Fetch(r.Url)
		if err!=nil{
			log.Printf("fetching url %s, %v",r.Url,err)
			continue
		}
		parseResult :=r.ParseFunc(body)
		requests = append(requests,parseResult.Requests...)

		for _,items := range parseResult.Items{
			log.Printf("Got item %v",items)
		}
	}

}

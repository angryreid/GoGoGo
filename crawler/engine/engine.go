package engine

import (
	"go-learing/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		log.Printf("Engine: fetching url %s\n", r.Url)
		requests = requests[1:]
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error in fetching url %s: %v\n", r.Url, err)
			continue
		}
		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Request...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %s", item)
		}
	}
}

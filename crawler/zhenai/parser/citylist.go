package parser

import (
	"fmt"
	"go-learing/crawler/engine"
	"log"
	"regexp"
)

const cityListRegExp = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParserResult {
	reg := regexp.MustCompile(cityListRegExp)
	matches := reg.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	limit := 2
	for _, match := range matches {
		result.Items = append(result.Items, string(match[2]))
		fmt.Printf("City: %s, URL: %s\n", match[2], match[1])
		// match[2] city name
		// match[1] city URL
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(match[1]),
			ParserFunc: ParserCity,
		})
		limit--
		if limit == 0 {
			break
		}
	}
	log.Printf("cityList: fetched all city counts: %d\n", len(matches))
	return result
}

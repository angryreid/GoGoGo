package parser

import (
	"go-learing/crawler/engine"
	"log"
	"regexp"
)

const cityListRegExp = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParserResult {
	reg := regexp.MustCompile(cityListRegExp)
	matches := reg.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, match := range matches {
		//fmt.Printf("City: %s, URL: %s\n", match[2], match[1])
		result.Request = append(result.Request, engine.Request{
			Url:        string(match[1]),
			ParserFunc: engine.NilParser,
		})
	}
	log.Printf("fetched all city counts: %d\n", len(matches))
	return result
}

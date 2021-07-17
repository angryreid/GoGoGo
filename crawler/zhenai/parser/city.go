package parser

import (
	"fmt"
	"go-learing/crawler/engine"
	"log"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[\d]+)"[^>]*>([^<]+)</a>`)
)

func ParserCity(contents []byte) engine.ParserResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, match := range matches {
		curUserName := string(match[2])
		result.Items = append(result.Items, curUserName)
		fmt.Printf("User: %s \n", match[2])
		// match[2] city name
		// match[1] city URL
		result.Requests = append(result.Requests, engine.Request{
			Url: string(match[1]),
			ParserFunc: func(c []byte) engine.ParserResult {
				return ParseProfile(c, curUserName)
			},
		})
	}
	log.Printf("cityList: fetched current city users: %d\n", len(matches))
	return result
}

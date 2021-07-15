package main

import (
	"go-learing/crawler/fetcher"
)

func main() {

	bytes, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	ParserCityList(bytes)

}

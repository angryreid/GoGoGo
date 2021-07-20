package book

import (
	"fmt"
	"go-learing/crawler/engine"
)

func ParserBook(contents []byte) engine.ParserResult {
	result := engine.ParserResult{}
	fmt.Printf("%s", contents)
	return result
}

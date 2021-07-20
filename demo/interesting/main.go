package main

import (
	"go-learing/crawler/engine"
	book "go-learing/demo/interesting/parser"
)

func main() {
	engine.Run(engine.Request{Url: "https://book.douban.com/top250", ParserFunc: book.ParserBook})
}

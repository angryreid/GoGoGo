package main

import (
	"fmt"
	"go-learing/mid/interface/retriever/mock"
	"go-learing/mid/interface/retriever/real"

)

type  Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.dapenti.com")
}

func main() {
	var r Retriever
	r = mock.MRetriever{Contents: "google.com"}
	r = real.Retriever{}
	fmt.Println(download(r))
}

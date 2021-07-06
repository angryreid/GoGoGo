package main

import (
	"fmt"
	"go-learing/mid/interface/retriever/mock"
	"go-learing/mid/interface/retriever/real"
	"time"
)

type  Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.dapenti.com")
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch typeR := r.(type) {
	case mock.MRetriever:
		fmt.Println("Contents:", typeR.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", typeR.UserAgent)
	}
}

func main() {
	var r Retriever
	r = mock.MRetriever{Contents: "google.com"}
	inspect(r)

	// Type Assertion
	mockRetriever := r.(mock.MRetriever)
	fmt.Println("mockRetriever:", mockRetriever.Contents)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)
	//fmt.Println(download(r))

	// Type Assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println("realRetriever", realRetriever.UserAgent)


}

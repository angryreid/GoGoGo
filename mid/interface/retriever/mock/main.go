package mock

import "fmt"

type MRetriever struct {
	Contents string
}

func (r MRetriever) String() string {
	//panic("implement me")
	return fmt.Sprintf("MRetriever: {Contents=%s}", r.Contents)
}

func (r MRetriever) Get(url string) string  {
	return r.Contents
}
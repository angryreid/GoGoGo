package mock

type MRetriever struct {
	Contents string
}

func (r MRetriever) Get(url string) string  {
	return r.Contents
}
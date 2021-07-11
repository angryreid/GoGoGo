package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	const site = "http://www.dapenti.com"
	request, err := http.NewRequest(http.MethodGet, site, nil)

	if err != nil {
		panic(err)
	}

	client := http.Client{
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request) error {
			fmt.Println("Redirect: ", req)
			return nil
		},
	}

	//res, err := http.DefaultClient.Do(request)
	res, err := client.Do(request)

	defer res.Body.Close()

	s, err := httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)
}

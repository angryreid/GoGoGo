package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte) {
	reg := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
	matches := reg.FindAllSubmatch(contents, -1)
	for _, match := range matches {
		fmt.Printf("City: %s, URL: %s\n", match[2], match[1])
	}
	fmt.Println(len(matches))
}

func main() {
	res, err := http.Get("https://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Http Status Error")
		return
	}

	e := determineEncoding(res.Body)

	reader := transform.NewReader(res.Body, e.NewDecoder())

	all, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("    %s\n", all)
	printCityList(all)

}

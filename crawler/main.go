package main

import (
	"bufio"
	"encoding"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func determineEncoding(r io.Reader) encoding.BinaryMarshaler {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}

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

	all, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", all)

}

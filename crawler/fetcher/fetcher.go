package fetcher

import (
	"bufio"
	"fmt"
	"go-learing/crawler/model"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetch error: %v\n", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func Fetch(url string) ([]byte, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
		return nil, err
	}
	req.Header.Add("cookie", model.TempCookie)
	req.Header.Set("User-Agent", model.UserAgent) // to resolve 403

	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		//return nil, errors.New("http status error")
		return nil,
			fmt.Errorf("http status code error: %d\n", resp.StatusCode)
	}
	defer resp.Body.Close()

	bfReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bfReader)

	reader := transform.NewReader(bfReader, e.NewDecoder())

	return ioutil.ReadAll(reader)
}

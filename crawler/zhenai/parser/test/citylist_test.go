package parser

import (
	"go-learing/crawler/zhenai/parser"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contents, _ := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", contents)

	result := parser.ParserCityList(contents)

	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	//expectedCities := []string{
	//	"City 阿坝",
	//	"City 阿克苏",
	//	"City 阿拉善盟",
	//}

	if len(result.Requests) != resultSize {
		t.Errorf("Result should have %d requests; but had %d",
			resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url %d: %s; but got %s",
				i, url, result.Requests[i].Url)
		}
	}
	//if len(result.Items) != resultSize {
	//	t.Errorf("Result should have %d items; but had %d",
	//		resultSize, len(result.Items))
	//}
	//for i, city := range expectedCities {
	//	if result.Items[i].(string) != city {
	//		t.Errorf("expected city %d: %s; but got %s",
	//			i, city, result.Items[i])
	//	}
	//}
}

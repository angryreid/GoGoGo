package parser

import (
	"fmt"
	"go-learing/crawler/fetcher"
	"go-learing/crawler/model"
	"go-learing/crawler/zhenai/parser"
	"testing"
)

const profileTest = "https://album.zhenai.com/u/1539752189"

func TestParseProfile(t *testing.T) {
	contents, err := fetcher.Fetch(profileTest)
	//contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", contents)

	result := parser.ParseProfile(contents, "风中的蒲公英")
	fmt.Printf("Testing Profile: %v\n", result)
	expected := model.Profile{
		Name:       "风中的蒲公英",
		Gender:     "女",
		Marriage:   "离异",
		Age:        41,
		Height:     158,
		Weight:     48,
		Income:     "3001-5000元",
		Education:  "中专",
		Occupation: "公务员",
		HuKou:      "四川阿坝",
		Xingzuo:    "处女座",
		House:      "已购房",
		Car:        "未购车",
	}

	if len(result.Items) != 1 {
		t.Errorf("Items should have 1 item, but got %v", result.Items)
	} else {
		actual := result.Items[0]
		if actual != expected {
			t.Errorf("Expected %+v, but got %+v", expected, actual)
		}
	}
}
